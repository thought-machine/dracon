import argparse
import logging
from datetime import datetime
from api.proto import engine_pb2
from consumers.consumer import Consumer
from third_party.python.defectdojo_api import defectdojo

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

class DefectDojoConsumer(Consumer):

    def __init__(self, config: dict):
        self.processed_records = 0
        self.pvc_location = config.pvc_location
        self.api_key = config.api_key
        self.dojo_url = config.dojo_url
        self.dojo_user = str(config.dojo_user)
        self.dojo_product = config.dojo_product
        self.dojo_engagement = config.dojo_engagement
        self.dojo_user_id = config.dojo_user_id
        self.dojo_test_id = None
        self.dd = defectdojo.DefectDojoAPI(
            self.dojo_url, self.api_key, self.dojo_user, debug=False)

        if (self.pvc_location is None):
            raise AttributeError("PVC claim location is missing")

    def load_results(self) -> (list, bool):
        try:
            return self._load_enriched_results(), False
        except SyntaxError:
            return self._load_plain_results(), True

    def _load_plain_results(self):
        scan_results = engine_pb2.LaunchToolResponse()
        return self.load_files(scan_results, self.pvc_location)

    def _load_enriched_results(self):
        """Load a set of LaunchToolResponse protobufs into a list for processing"""
        return super().load_results()

    def _send_to_dojo(self, data: dict, dojo_test_id: int, start_date: str):

        severity_map = {0: "Low", 1: "Low",
                        2: "Medium", 3: "High", 4: "Critical", 5: 'Info'}
        logger.debug("Sending to dojo")
        # todo (spyros): it also support marking findings as duplicates, if first_found is in
        #  the past this can be a duplicate
        impact = "Possible product vulnerability"
        active = True
        verified = False
        mitigation = "Please triage and resolve"
        self.processed_records += 1
        description = ("scan_id: %s \n tool_name: %s \n type: %s \n confidence: %s\n"
                       "original_path=%s \n original description: %s" % (data['scan_id'],
                                                                         data['tool_name'],
                                                                         data['type'],
                                                                         data['confidence'],
                                                                         data['target'],
                                                                         data['description']))

        finding = self.dd.create_finding(data['title'],
                                         description,
                                         severity_map[data['severity']],
                                         0,
                                         start_date,
                                         self.dojo_product,
                                         self.dojo_engagement,
                                         dojo_test_id,
                                         self.dojo_user_id,
                                         impact,
                                         active,
                                         verified,
                                         mitigation,
                                         references=None,
                                         build=None,
                                         line=0,
                                         # TODO (spyros): this is a hack so we
                                         #  can mark issues as "viewed",
                                         #  remove when
                                         # https://github.com/DefectDojo/django-DefectDojo/issues/1609
                                         # gets implemented
                                         under_review=True,
                                         file_path=data['target'],
                                         false_p=str(data['false_positive']))
        if not finding.success:
            raise Exception(
                "Couldn't communicate to DefectDojo error message: %s" % finding.message)
        

    def send_results(self, collected_results: list, raw_issue: bool):
        """
        Take a list of *ToolResponse protobufs and sends them to DefectDojo
        If results are enriched, only the new, non-false positive results will be sent
        :param collected_results: list of LaunchToolResponse protobufs
        """
        for sc in collected_results:
            logger.debug("handling result")
            for iss in sc.issues:
                logger.debug("handling issue")
                if raw_issue:
                    logger.debug("issue is raw")
                    scan = sc
                    issue = iss
                    first_found = scan.scan_info.scan_start_time.ToJsonString()
                    false_positive = False
                else:
                    logger.debug("issue %s is enriched!" % iss.raw_issue.title)
                    issue = iss.raw_issue
                    first_found = iss.first_seen.ToJsonString()
                    false_positive = iss.false_positive
                    scan = sc.original_results
                    if iss.count > 1:
                        logger.debug('Issue %s is a duplicate, count= %s, skipping' %
                                     (issue.title, iss.count))
                        continue
                    if false_positive:
                        logger.debug(
                            'Issue %s has been marked as a false positive, skipping' % issue.title)
                        continue

                data = {
                    'scan_start_time': scan.scan_info.scan_start_time.ToJsonString(),
                    'scan_id': scan.scan_info.scan_uuid,
                    'tool_name': scan.tool_name,
                    'target': issue.target,
                    'type': issue.type,
                    'title': issue.title,
                    'severity': issue.severity,
                    'cvss': issue.cvss,
                    'confidence': issue.confidence,
                    'description': issue.description,
                    'first_found': first_found,
                    'false_positive': false_positive
                }
                start_date = datetime.strptime(
                    data.get('scan_start_time'), '%Y-%m-%dT%H:%M:%SZ').date().isoformat()
                if not self.dojo_test_id:
                    logger.info("Test %s doesn't exist, creating" %
                                scan.scan_info.scan_uuid)
                    start_date = datetime.strptime(
                        data.get('scan_start_time'), '%Y-%m-%dT%H:%M:%SZ').date().isoformat()
                    end_date = datetime.utcnow().date()
                    test_type = 2  # static Check sounds most generic, the python client
                    # won't accept adding custom title
                    # TODO (spyros): commit upstream
                    environment = 1  # development
                    test = self.dd.create_test(self.dojo_engagement,
                                               str(test_type),
                                               str(environment),
                                               start_date,
                                               end_date.isoformat())
                    if not test.success:
                        raise Exception(
                            "Couldn't create defecto dojo test: %s" % test.message)

                    self.dojo_test_id = test.id()
                self._send_to_dojo(data, self.dojo_test_id, start_date)


def main():
    try:
        parser = argparse.ArgumentParser()
        parser.add_argument(
            '--pvc_location', help='The location of the scan results')
        parser.add_argument(
            '--raw', help='if it should process raw or enriched results', action="store_true")
        parser.add_argument(
            '--api_key', help='the api key for the defect dojo instance to connect to')
        parser.add_argument('--dojo_url', help='defectdojo api target url')
        parser.add_argument('--dojo_user', help='defectdojo user')
        parser.add_argument(
            '--dojo_product', help='defectdojo product for which the findings')
        parser.add_argument(
            '--dojo_engagement', help='defectdojo ci/cd style engagment for which you want to add'
            ' the test and findings')
        parser.add_argument(
            '--dojo_user_id', help='defectdojo id for the user you just specified')
        args = parser.parse_args()
        dd = DefectDojoConsumer(args)
    except AttributeError as e:
        raise Exception('A required argument is missing: ' + str(e))

    logger.info('Loading results from %s' % str(dd.pvc_location))
    collected_results, raw = dd.load_results()
    if len(collected_results) == 0:
        raise Exception('Unable to load results from the filesystem')

    logger.info("gathered %s results" % len(collected_results))
    logger.info("Reading raw: %s " % raw)
    dd.send_results(collected_results, raw)
    logger.info('Done, processed %s records!' % dd.processed_records)


if __name__ == '__main__':
    main()
