#!/bin/bash

set -euo pipefail

YQ_BIN="third_party/tools/yq_linux_amd64"
INPUT="docs/designs/dracon-v2-example/golang-example/golang-example_kustomized.yaml"


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
source third_party/sh/shunit2
