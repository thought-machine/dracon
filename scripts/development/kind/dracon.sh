#!/bin/bash
# This script is a wrapper around dracon commands to give a 
# development experience as close to the release experience as possible.
# e.g. a released command could be `dracon setup --namespace dracon --pipeline examples/pipelines/python-project`
# in development, we provide the ability to run this prefixed with `plz`, 
# e.g. `plz dracon setup --namespace dracon --pipeline examples/pipelines/python-project`

kubernetes_context=$(kubectl config current-context)
if [ "${kubernetes_context}" != "kind-dracon" ]; then
  util::prompt "Are you sure you would like to run dracon against '${kubernetes_context}'?"
fi

kind_cluster_name="${kubernetes_context//kind-/}"

args=()
while [ $# -gt 0 ]; do
    case $1 in
        --pipeline)
            pipeline_target="//$2:dev"
            if [[ "${args[@]}" == *"run "* ]]; then
                util::load_development_images_into_kind "${kind_cluster_name}" "$pipeline_target"
            fi

            # convert pipeline flag to generated pipeline which has the local docker image tags
            ./pleasew build "$pipeline_target"
            pipeline_out_dir="$(dirname $(./pleasew query output $pipeline_target))"
            args+=("$1" "$pipeline_out_dir")
            shift 2
        ;;
        --*)
            args+=("$1" "$2")
            shift 2
        ;;
        *)
            args+=("$1")
            shift 1
        ;;
    esac
done

"$DRACON_BIN" "${args[@]}"
