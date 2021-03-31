#!/bin/bash

set -euo pipefail

# util::load_development_images_into_kind builds and loads the given target's Docker images into the KinD cluster.
util::load_development_images_into_kind() {
    local kind_cluster_name pipeline_target
    kind_cluster_name="$1"
    pipeline_target="$2"

    docker_fqn_targets=($(./pleasew query deps \
        --include docker-fqn \
        "$pipeline_target" \
        | sort -u \
        | awk '{ print $1 }'
    ))
    if [ "${#docker_fqn_targets[@]}" == 0 ]; then
        return
    fi
    util::rinfor "building ${#docker_fqn_targets[@]} image(s)"
    docker_load_targets=($(printf '%s\n' "${docker_fqn_targets[@]}" | sed 's/_fqn$/_load/g'))
    ./pleasew -p -v 2 --colour run parallel --quiet "${docker_load_targets[@]}"
    for docker_fqn_target in "${docker_fqn_targets[@]}"; do
        util::rinfor "loading $docker_fqn_target into kind cluster '$kind_cluster_name'"
        docker_fqn=$(<$(./pleasew query output $docker_fqn_target))
        "$KIND_BIN" load docker-image \
            "${docker_fqn}" \
            --name "${kind_cluster_name}" > /dev/null
    done
    util::rsuccess "loaded development images in to kind cluster '$kind_cluster_name'"
}
