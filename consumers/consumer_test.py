import unittest

from google.protobuf.timestamp_pb2 import Timestamp

from consumers.consumer import Consumer
from gen import engine_pb2
from gen import issue_pb2


class ExampleConsumer(Consumer):
    '''
    Example implementation of a consumer so that we can instantiate an object
    that implements the Consumer class
    '''

    def __init__(self, config):
        super().__init__(config)

    def load_results(self):
        return super().load_results()

    def send_results(self, collected_results):
        super().send_results(collected_results)


class TestConsumer(unittest.TestCase):

    def setUp(self):
        self.config = {
            'dry_run': True,
            'es_index': 'dracon',
            'es_url': 'https://some_test.url.somewhere.io:443',
            'pvc_location': './'
        }

        # Create an scan results object and serialize it to a file
        ts = Timestamp()
        ts.FromJsonString("1991-01-01T00:00:00Z")
        scan_results = engine_pb2.LaunchToolResponse(
            scan_info=engine_pb2.ScanInfo(
                scan_uuid='dd1794f2-544d-456b-a45a-a2bec53633b1',
                scan_start_time=ts,
            ),
            tool_name='bandit',
        )

        issue = issue_pb2.Issue()
        issue.target = 'target.py:0'
        scan_results.issues.extend([issue])

        enriched_scan_results = engine_pb2.EnrichedLaunchToolResponse(
            original_results=scan_results,
        )

        f = open(self.config['pvc_location'] + "example_response.pb", "wb")
        scan_proto_string = enriched_scan_results.SerializeToString()
        f.write(scan_proto_string)
        f.close()

    def test_correct_proto(self):
        '''A basic test to check we can take a configuration and correctly read from it'''

        consumer = ExampleConsumer(self.config)
        collected_results = consumer.load_results()

        scan_result = collected_results.pop()
        raw_scan_result = scan_result.original_results
        self.assertEqual(
            raw_scan_result.scan_info.scan_uuid,
            'dd1794f2-544d-456b-a45a-a2bec53633b1'
        )
        self.assertEqual(raw_scan_result.tool_name, 'bandit')
        self.assertEqual(
            raw_scan_result.scan_info.scan_start_time.ToJsonString(),
            '1991-01-01T00:00:00Z'
        )

        issue = raw_scan_result.issues.pop()
        self.assertEqual(issue.target, 'target.py:0')


if __name__ == '__main__':
    unittest.main()
