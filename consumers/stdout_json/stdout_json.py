import sys
import json
from gen import engine_pb2
from consumers.consumer import Consumer
import logging
from utils.file_utils import load_files
import argparse


logger = logging.getLogger(__name__)


class StdoutJsonConsumer(Consumer):

    def __init__(self, config: dict):
        print("Starting Consumer")
        self.pvc_location = config.pvc_location
        print("Reading from %s" % self.pvc_location)

        if (self.pvc_location is None):
            raise AttributeError("PVC claim location is missing")

    def load_results(self) -> (list, bool):
        try:
            return self._load_enriched_results(), False
        except SyntaxError:
            return self._load_plain_results(), True

    def _load_plain_results(self):
        scan_results = engine_pb2.LaunchToolResponse()
        return load_files(scan_results, self.pvc_location)

    def _load_enriched_results(self):
        """Load a set of LaunchToolResponse protobufs into a list for processing"""
        return super().load_results()

    def print_data(self, data):
        print(json.dumps(data))

    def send_results(self, collected_results: list, raw: bool):
        """
        Take a list of LaunchToolResponse protobufs and sends them to Elasticsearch

        :param collected_results: list of LaunchToolResponse protobufs
        """
        for sc in collected_results:
            for iss in sc.issues:
                if raw:
                    scan = sc
                    issue = iss
                    first_found = scan.scan_info.scan_start_time.ToJsonString()
                    false_positive = False
                else:
                    issue = iss.raw_issue
                    first_found = iss.first_seen.ToJsonString()
                    false_positive = iss.false_positive
                    scan = sc.original_results
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
                self.print_data(data)


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument(
        '--pvc_location', help='The location of the scan results')
    parser.add_argument(
        '--raw', help='if it should process raw or enriched results', action="store_true")
    args = parser.parse_args()
    ec = StdoutJsonConsumer(args)
    try:
        print('Loading results from %s' % str(ec.pvc_location))
        collected_results, raw = ec.load_results()
        print("gathered %s results"%len(collected_results))
        print("Reading raw: %s "%len(raw))
    except SyntaxError as e:
        logger.error('Unable to load results from %s: ' % str(e))
        sys.exit(-1)

    ec.send_results(collected_results, args.raw)
    print('Done!')


if __name__ == '__main__':
    print("Consumer Stdout JSON running")
    main()
