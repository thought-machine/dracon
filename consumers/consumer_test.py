import shutil
import tempfile
import unittest

from api.proto import engine_pb2, issue_pb2
from consumers.consumer import Consumer
from third_party.python.google.protobuf.timestamp_pb2 import Timestamp


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

        scan_results = engine_pb2.LaunchToolResponse(
            scan_info=engine_pb2.ScanInfo(
                scan_uuid='dd1794f2-544d-456b-a45a-a2bec53633b1',
            ),
            tool_name='bandit',
        )
        self.tmp_root_dir = tempfile.mkdtemp()
        _, self.tmpfile = tempfile.mkstemp(
            suffix=".pb", prefix="example_response_", dir=self.tmp_root_dir)
        with open(self.tmpfile, "wb") as f:
            serialized_proto = scan_results.SerializeToString()
            f.write(serialized_proto)

        # Duplicate the serialized protobuf into a subfolder to check recursion
        self.tmp_subdir = tempfile.mkdtemp(dir=self.tmp_root_dir)
        _, self.tmpfile2 = tempfile.mkstemp(
            suffix=".pb", prefix="example_response_copy_", dir=self.tmp_subdir)
        with open(self.tmpfile2, "wb") as f:
            serialized_proto = scan_results.SerializeToString()
            f.write(serialized_proto)

        # Create a malformed protobuf to check we handle it gracefully
        malformed_proto = serialized_proto[10:]
        _, self.malformed = tempfile.mkstemp(
            suffix=".pb", prefix="malformed_", dir=self.tmp_root_dir)
        with open(self.malformed, "wb") as f:
            f.write(malformed_proto)

        print(self.tmp_root_dir, self.tmp_subdir,
              self.tmpfile, self.tmpfile2, self.malformed)

    def tearDown(self):
        shutil.rmtree(self.tmp_root_dir)

    def test_load_filecorrect_proto(self):
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

    def test_load_file_proto_read(self):
        '''Test we can load protos and read from them correctly
           Also ensures we handled malformed protobufs gracefully
        '''
        scan_result_proto = engine_pb2.LaunchToolResponse()
        cons = ExampleConsumer(self.config)
        collected_results = cons.load_files(
            scan_result_proto, self.tmp_root_dir)
        result = collected_results.pop()

        self.assertEqual(result.scan_info.scan_uuid,
                         'dd1794f2-544d-456b-a45a-a2bec53633b1')

    def test_load_file_search(self):
        '''Check that the recursive file detection is working as expected'''
        scan_result_proto = engine_pb2.LaunchToolResponse()
        cons = ExampleConsumer(self.config)
        collected_results = cons.load_files(
            scan_result_proto, self.tmp_root_dir)
        self.assertEqual(len(collected_results), 2)


if __name__ == '__main__':
    unittest.main()
