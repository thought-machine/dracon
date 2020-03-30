import shutil
import tempfile
import unittest
import logging
from unittest import mock
from collections import namedtuple

from consumers.defectdojo_c import defectdojo_c
from api.proto import engine_pb2, issue_pb2

from third_party.python.google.protobuf.timestamp_pb2 import Timestamp

logger = logging.getLogger(__name__)

CreateTestRet = namedtuple('CreateTestRet', ['id', 'success'])


class ConsumerMockConfig:

    def __init__(self):
        self.pvc_location = '/tmp/'


class DefectDojoConsumerTest(unittest.TestCase):

    def setUp(self):
        self.dojo_url = 'http://dojo.local/'
        self.dojo_api_key = ''
        self.config = ConsumerMockConfig()
        self.config.dojo_url = self.dojo_url
        self.config.api_key = self.dojo_api_key
        self.config.dojo_user = 'testuser'
        self.config.dojo_user_id = '1'
        self.config.dojo_product = 1
        self.config.dojo_engagement = 1
        self.config.raw = False

        scan_start_time = Timestamp()
        scan_start_time.FromJsonString("1991-01-01T00:00:00Z")
        scan_info = engine_pb2.ScanInfo(
            scan_start_time=scan_start_time,
            scan_uuid='dd1794f2-544d-456b-a45a-a2bec53633b1'
        )
        scan_results = engine_pb2.LaunchToolResponse(
            scan_info=scan_info
        )
        scan_results.tool_name = 'unit_tests'

        #  Raw results
        issue = issue_pb2.Issue()
        issue.target = 'target.py:0'
        issue.type = "test"
        issue.title = "test title"
        issue.cvss = 2.0
        issue.description = "test.description"
        issue.severity = issue_pb2.Severity.SEVERITY_LOW
        issue.confidence = issue_pb2.Confidence.CONFIDENCE_LOW
        scan_results.issues.extend([issue])

        # Enriched, duplicate and False Positive results
        first_seen = Timestamp()
        first_seen.FromJsonString("1992-02-02T00:00:00Z")
        enriched_issue = issue_pb2.EnrichedIssue(first_seen=first_seen)
        enriched_issue.raw_issue.CopyFrom(issue)
        enriched_issue.count = 2
        enriched_issue.false_positive = True

        enriched_scan_results = engine_pb2.EnrichedLaunchToolResponse(
            original_results=scan_results,
        )
        enriched_scan_results.issues.extend([enriched_issue])

        # Enriched, unique, false positive result
        enriched_issue.count = 0
        enriched_issue.false_positive = True
        issue.target = 'target0.py:0'
        issue.type = "test0"
        issue.title = "test0 title0"
        enriched_issue.raw_issue.CopyFrom(issue)
        enriched_scan_results.issues.extend([enriched_issue])

        # Enriched, unique, true positive result
        enriched_scan_results.issues.extend([enriched_issue])
        enriched_issue.count = 0
        enriched_issue.false_positive = False
        issue.target = 'target1.py:0'
        issue.type = "test1"
        issue.title = "test1 title1"
        enriched_issue.raw_issue.CopyFrom(issue)
        enriched_scan_results.issues.extend([enriched_issue])
        self.enriched_dtemp = tempfile.mkdtemp(
            prefix="enriched_", dir=self.config.pvc_location)
        self.enriched, _ = tempfile.mkstemp(
            prefix="enriched_", dir=self.enriched_dtemp, suffix=".pb")

        self.raw_dtemp = tempfile.mkdtemp(
            prefix="raw_", dir=self.config.pvc_location)
        self.raw, _ = tempfile.mkstemp(
            prefix="raw_", dir=self.raw_dtemp, suffix=".pb")

        f = open(self.enriched, "wb")
        scan_proto_string = enriched_scan_results.SerializeToString()
        f.write(scan_proto_string)
        f.close()

        f = open(self.raw, "wb")
        scan_proto_string = scan_results.SerializeToString()
        f.write(scan_proto_string)
        f.close()

    def tearDown(self):
        shutil.rmtree(self.raw_dtemp)
        shutil.rmtree(self.enriched_dtemp)

    def _create_consumer(self):
        return defectdojo_c.DefectDojoConsumer(self.config)

    @mock.patch('third_party.python.'
                'defectdojo_api.defectdojo.DefectDojoAPI.create_test')
    @mock.patch('third_party.python.'
                'defectdojo_api.defectdojo.DefectDojoAPI.create_finding')
    def test_send_results(self, mock_create_finding, mock_create_test):
        data = {
            'scan_start_time': "1991-01-01T00:00:00Z",
            'scan_id': 'dd1794f2-544d-456b-a45a-a2bec53633b1',
            'tool_name': 'unit_tests',
            'target': 'target.py:0',
            'type': "test",
            'title': "test title",
            'severity': 1,
            'cvss': 2.0,
            'confidence': 1,
            'description': "test.description",
            'first_found': '',
            'false_positive': 'False'
        }

        mock_create_test.return_value = CreateTestRet(id=lambda: 1, success=True)
        self.config.pvc_location = self.raw_dtemp
        consumer = defectdojo_c.DefectDojoConsumer(self.config)
        results, _ = consumer.load_results()

        impact = "Possible product vulnerability"
        active = True
        verified = False
        mitigation = "Please triage and resolve"

        description = ("scan_id: %s \n tool_name: %s \n type: %s \n confidence: %s\n"
                       "original_path=%s \n original description: %s" % (data['scan_id'],
                                                                         data['tool_name'],
                                                                         data['type'],
                                                                         data['confidence'],
                                                                         data['target'],
                                                                         data['description']))

        # Test correct handling of Raw results
        consumer.send_results(results, True)
        mock_create_finding.assert_called_with(data['title'],
                                               description,
                                               'Low',
                                               0,
                                               '1991-01-01',
                                               self.config.dojo_product,
                                               self.config.dojo_engagement,
                                               1,
                                               self.config.dojo_user_id,
                                               impact,
                                               active,
                                               verified,
                                               mitigation,
                                               references=None,
                                               build=None,
                                               line=0,
                                               file_path=data['target'],
                                               false_p=str(data['false_positive']),
                                               under_review=True)

        # test correct handling of enriched, False Positives
        self.config.pvc_location = self.enriched_dtemp
        consumer = defectdojo_c.DefectDojoConsumer(self.config)
        results, raw = consumer.load_results()

        data['first_found'] = "1992-02-02T00:00:00Z"
        data['false_positive'] = False
        data['type'] = 'test1'
        data['target'] = 'target1.py:0'

        description = ("scan_id: %s \n tool_name: %s \n type: %s \n confidence: %s\n"
                       "original_path=%s \n original description: %s" % (data['scan_id'],
                                                                         data['tool_name'],
                                                                         data['type'],
                                                                         data['confidence'],
                                                                         data['target'],
                                                                         data['description']))

        consumer.send_results(results, False)
        mock_create_finding.assert_called_with('test1 title1',
                                               description,
                                               'Low',
                                               0,
                                               '1991-01-01',
                                               self.config.dojo_product,
                                               self.config.dojo_engagement,
                                               1,
                                               self.config.dojo_user_id,
                                               impact,
                                               active,
                                               verified,
                                               mitigation,
                                               references=None,
                                               build=None,
                                               line=0,
                                               file_path='target1.py:0',
                                               false_p=str(data['false_positive']),
                                               under_review=True)


if __name__ == '__main__':
    unittest.main()
