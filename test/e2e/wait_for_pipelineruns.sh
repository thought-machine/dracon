#!/bin/bash
# This script waits for pipelineruns to finish and exits depending on if all pipelineruns were successful or not.
set -euo pipefail


function all_pipelineruns_succeeded {
    pipeline_runs=($(kubectl -n dracon get pipelineruns.tekton.dev -oyaml | \
        $YQ_BIN e -N '.items[] | .metadata.name + "," + .status.conditions[0].reason' -))
    has_pipeline_running=false
    has_pipeline_error=false

    for pipeline_run in "${pipeline_runs[@]}"; do
        name="$(echo "$pipeline_run" | cut -f1 -d,)"
        status="$(echo "$pipeline_run" | cut -f2 -d,)"

        case "$status" in
            "Running")
                util::rinfor "$name is Running..."
                has_pipeline_running=true
            ;;
            "Succeeded")
            ;;
            *)
                util::error "$name is $status"
                has_pipeline_error=true
        esac
    done

    if [ "$has_pipeline_error" = true ]; then
        return 1
    fi

    if [ "$has_pipeline_running" = true ]; then
        return 2
    fi

    return 0
}

time_limit_secs=3600
intervals=100
sleep_interval=$(($time_limit_secs/$intervals))
attempts=0

util::info "waiting for all pipelines to succeed within ${time_limit_secs}s"

set +x
until all_pipelineruns_succeeded; do
    ec="$?"
    if [ "$ec" == 1 ]; then
        exit 1
    fi
    if [ $attempts -eq $intervals ]; then
        util::error "timed out"
        exit 1
    fi
    attempts=$((attempts + 1))
    sleep $sleep_interval
done
set -e

util::rsuccess "all pipelineruns completed successfully"
