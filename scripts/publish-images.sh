#!/bin/bash
set -eo pipefail

if [ -n "${DOCKERHUB_USERNAME}" ]; then
  echo "${DOCKERHUB_PASSWORD}" | docker login --username "${DOCKERHUB_USERNAME}" --password-stdin
fi

version=$(git describe --always)

docker_rules=$(./pleasew query alltargets --include docker-build //...)

fqns_to_push=()

export DOCKER_BUILDKIT=1
for rule in ${docker_rules}; do
  ./pleasew run "${rule}_load"
  fqn=$(cat $(./pleasew build "${rule}_fqn" | tail -n1 | tr -s " "))
  repo="$(echo "${fqn}" | cut -f1 -d\:)"
  fqn_version="${repo}:${version}"
  echo ""
  echo "-> tagging as ${fqn_version}"
  docker tag "${fqn}" "${fqn_version}"
  fqns_to_push+=("${fqn_version}")
done

for fqn_to_push in "${fqns_to_push[@]}"; do
  echo "-> pushing as ${fqn_to_push}"
  docker push "${fqn_to_push}"
done
