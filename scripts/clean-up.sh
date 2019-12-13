#!/bin/bash

function deleteAll() {
  local resource=$1
  kubectl get $resource -l project=dracon -o custom-columns=NAME:.metadata.name | tail -n+2 | xargs kubectl delete $resource
}

deleteAll "pipelineruns.tekton.dev"
deleteAll "pipelineresources.tekton.dev"
deleteAll "taskruns.tekton.dev"
deleteAll "pipelines.tekton.dev"
deleteAll "tasks.tekton.dev"
