#!/usr/bin/env python3

import tempfile

import unittest

from producers.python_bandit.bandit import parse_tool_results
import producers.producer as shared
from gen import engine_pb2
from gen import issue_pb2


class TestBanditProducer(unittest.TestCase):

    tool_out = b"""
{
    "results": [
        {
            "code": "4         self.output =",
            "filename": "./foo/bar.py",
            "issue_confidence": "MEDIUM",
            "issue_severity": "MEDIUM",
            "issue_text": "Probable insecure usage of temp file/directory.",
            "line_number": 4,
            "line_range": [4, 5],
            "more_info": "https://bandit.readthedocs.io/en/latest/plugins/b108_hardcoded_tmp_directory.html",
            "test_id": "B108",
            "test_name": "hardcoded_tmp_directory"
        }
    ]
}
"""

    def test_bandit_produce(self):
        with tempfile.NamedTemporaryFile() as i_f, \
                tempfile.NamedTemporaryFile() as o_f:
            i_f.write(self.tool_out)
            i_f.flush()
            args = shared.parse_flags(['-out', o_f.name, '-in', i_f.name])
            tool_results = shared.parse_in_file_json(args)
            issues = parse_tool_results(tool_results)
            shared.write_dracon_out(args, "bandit", issues)

            res = o_f.read()
            pblaunch = engine_pb2.LaunchToolResponse()
            pblaunch.ParseFromString(res)
            self.assertEqual(engine_pb2.LaunchToolResponse(
                tool_name="bandit",
                issues=[issue_pb2.Issue(
                    target="./foo/bar.py:[4, 5]",
                    type="hardcoded_tmp_directory",
                    title="hardcoded_tmp_directory",
                    severity="SEVERITY_MEDIUM",
                    cvss=0,
                    confidence="CONFIDENCE_MEDIUM",
                    description="Probable insecure usage of temp file/directory.",
                    source="unknown"

                )]
            ), pblaunch)


if __name__ == '__main__':
    unittest.main()
