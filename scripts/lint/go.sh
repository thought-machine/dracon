#!/bin/bash
set -Eeuo pipefail

util::infor "linting go files"

dirs=($(./pleasew query alltargets --include=go | grep -v third_party | cut -f1 -d":" | cut -c 3- | sort -u))
if ! "${GO_LINT}" -set_exit_status ${dirs[@]}; then
  util::rerror "go files failed lint. To fix format errors, please run:
  $ ./pleasew run //scripts/fmt:go"
    exit 1
fi

util::rsuccess "linted go files"
