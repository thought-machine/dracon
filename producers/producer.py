#!/usr/bin/env python3

import os
import logging
import json
import argparse

from gen import engine_pb2
from gen import issue_pb2

logger = logging.getLogger(__name__)

parser = argparse.ArgumentParser()

__source_dir = "/dracon/source"


def parse_flags(args: object) -> object:
    """
    Parses the input flags for a producer
    """
    parser.add_argument('-in', help='tool results file')
    parser.add_argument('-out', help='producer output file')
    return parser.parse_args(args)


def parse_in_file_json(args: object) -> dict:
    """
    A generic method to return a tool's JSON results file as a dict
    """
    results_file = vars(args)['in']
    with open(results_file) as f:
        data = f.read()
        return json.loads(data)


def write_dracon_out(args: object, tool_name: str, issues: [issue_pb2.Issue]):
    """
    A method to write the resulting protobuf to the output file
    """
    out_file = vars(args)['out']
    source = __get_meta_source()
    clean_issues = []
    for iss in issues:
        iss.description = iss.description.replace(__source_dir, ".")
        iss.title = iss.title.replace(__source_dir, ".")
        iss.target = iss.target.replace(__source_dir, ".")
        iss.source = source
        clean_issues.append(iss)

    ltr = engine_pb2.LaunchToolResponse(
        tool_name=tool_name,
        issues=issues
    )

    with open(out_file, 'ab') as f:
        f.write(ltr.SerializeToString())


__meta_src_file = ".source.dracon"


def __get_meta_source() -> str:
    """
    This obtains the source address in the __meta_src_file from the source workspace
    """
    meta_src_path = os.path.join(__source_dir, __meta_src_file)
    if os.path.exists(meta_src_path):
        with open(meta_src_path) as f:
            return f.read().strip()
    return "unknown"
