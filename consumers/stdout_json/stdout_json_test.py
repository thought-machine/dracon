import unittest
import tempfile
import shutil

from google.protobuf.timestamp_pb2 import Timestamp
from unittest import mock
from consumers.stdout_json import stdout_json
from gen import engine_pb2
from gen import issue_pb2
import logging
from utils import test_utils

logger = logging.getLogger(__name__)


class TestJSONConsumer(unittest.TestCase):
    
    def setUp(self):    
        self.config = test_utils.ConsumerMockConfig()
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

        issue = issue_pb2.Issue()
        issue.target = 'target.py:0'
        issue.type = "test"
        issue.title = "test title"
        issue.cvss = 2.0
        issue.description = "test.description"
        issue.severity = issue_pb2.Severity.SEVERITY_LOW
        issue.confidence = issue_pb2.Confidence.CONFIDENCE_LOW

        scan_results.issues.extend([issue])
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
        return stdout_json.StdoutJsonConsumer(self.config)

    def test_load_results(self):
        self.config.pvc_location = self.enriched_dtemp
        consumer = stdout_json.StdoutJsonConsumer(self.config)
        _, raw = consumer.load_results()
        self.assertFalse(raw)

        self.config.pvc_location = self.raw_dtemp
        consumer = stdout_json.StdoutJsonConsumer(self.config)
        _, raw = consumer.load_results()
        self.assertTrue(raw)

    @mock.patch(
        "consumers.stdout_json.stdout_json"
        ".StdoutJsonConsumer"
        ".print_data")
    def test_send_results(self, mocked_print_data):
        mocked_print_data.return_value = ""

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
            'false_positive': ''
        }

        self.config.pvc_location = self.raw_dtemp
        consumer = stdout_json.StdoutJsonConsumer(self.config)
        results, raw = consumer.load_results()

        consumer.send_results(results, raw)
        data['first_found'] = data['scan_start_time']
        data['false_positive'] = False
        mocked_print_data.assert_called_with(data)

        self.config.pvc_location = self.enriched_dtemp
        consumer = stdout_json.StdoutJsonConsumer(self.config)
        results, raw = consumer.load_results()

        consumer.send_results(results, raw)
        data['first_found'] = "1992-02-02T00:00:00Z"
        data['false_positive'] = True
        mocked_print_data.assert_called_with(data)


if __name__ == '__main__':
    unittest.main()
