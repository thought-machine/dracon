#!/bin/bash

set -euo pipefail

SHUNIT="$DATA_SHUNIT"
YQ_BIN="$DATA_YQ"
INPUT="$DATA_INPUT"

testEquality() {
    assertEquals 1 1
}

testPipelineTasksMergeParamValues() {
    testTaskRef=$("$YQ_BIN" eval-all '.spec.tasks[].params[].value' "$INPUT" | grep "^-")
    expected=$(cat <<EOF
- \$(tasks.b-task.results.issues)
- \$(tasks.a-task.results.issues)
EOF
)
    assertEquals "${expected}" "${testTaskRef}"
}

testPipelineTaskRef() {
    taskRef=$("$YQ_BIN" eval-all '.spec.tasks[].taskRef.name' "$INPUT" | grep "a-task")
    assertEquals "a-task-test" "${taskRef}"
}

# Load shUnit2.
source "$SHUNIT"
