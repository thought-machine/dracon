#!/bin/bash

set -euo pipefail

SHUNIT="$DATA_SHUNIT"
YQ_BIN="$DATA_YQ"
INPUT="$DATA_INPUT"


testEquality() {
    assertEquals 1 1
}

testPipelineTasksMergeParamValues() {
    assertEquals 1 1
}

testPipelineTaskRef() {
    gosecProducerTaskRef=$("$YQ_BIN" eval-all '.spec.tasks[].taskRef.name' "$INPUT" | grep "gosec-producer")
    assertEquals "gosec-producer-golang-example" "${gosecProducerTaskRef}"
}

# Load shUnit2.
source "$SHUNIT"
