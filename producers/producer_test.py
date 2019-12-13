#!/usr/bin/env python3

import tempfile
import unittest

import producers.producer as shared

from gen import engine_pb2
from gen import issue_pb2


class TestProducer(unittest.TestCase):

    def test_write_dracon_out(self):
        with tempfile.NamedTemporaryFile() as fp:
            args = shared.parse_flags(['-out', fp.name, '-in', 'foo'])
            shared.write_dracon_out(args, "foo", [
                issue_pb2.Issue(
                    target="/dracon/source/foobar",
                    title="/dracon/source/barfoo",
                    description="/dracon/source/example.yaml",
                )
            ])
            fp.flush()
            res = fp.read()
            pblaunch = engine_pb2.LaunchToolResponse()
            pblaunch.ParseFromString(res)
            self.assertEqual(engine_pb2.LaunchToolResponse(
                tool_name="foo",
                issues=[issue_pb2.Issue(
                    target="./foobar",
                    title="./barfoo",
                    description="./example.yaml",
                    source="unknown"
                )]
            ), pblaunch)


if __name__ == '__main__':
    unittest.main()
