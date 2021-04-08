#!/bin/bash
set -Eeuo pipefail

util::infor "formatting BUILD files"
./pleasew fmt --write
util::rsuccess "formatted BUILD files"
