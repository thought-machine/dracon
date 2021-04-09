#!/bin/bash
set -Eeuo pipefail

util::infor "checking BUILD files"
if ! ./pleasew fmt --quiet; then
  util::rerror "BUILD files incorrectly formatted. Please run:
  $ ./pleasew run //scripts/fmt:plz"
  exit 1
fi
util::rsuccess "checked BUILD files"
