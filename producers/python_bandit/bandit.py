#!/usr/bin/env python3

import logging
import sys
import producers.producer as shared
from gen import issue_pb2


logger = logging.getLogger(__name__)


def parse_tool_results(results: dict) -> [issue_pb2.Issue]:
    issues = []
    for res in results['results']:
        iss = parse_issue(res)
        issues.append(iss)
    return issues


def parse_issue(rec_issue: dict) -> issue_pb2.Issue:
    return issue_pb2.Issue(
        target=f"{rec_issue['filename']}:{rec_issue['line_range']}",
        type=rec_issue['test_name'],
        title=rec_issue['test_name'],
        severity=f"SEVERITY_{rec_issue['issue_severity']}",
        cvss=0,
        confidence=f"CONFIDENCE_{rec_issue['issue_confidence']}",
        description=rec_issue['issue_text']
    )


if __name__ == "__main__":
    args = shared.parse_flags(sys.argv[1:])
    tool_results = shared.parse_in_file_json(args)
    issues = parse_tool_results(tool_results)
    shared.write_dracon_out(args, "bandit", issues)
