#!/bin/bash

# This script uploads Dracon container images to Docker Hub, as well as the GitHub Container
# Registry if the script is being executed by a GitHub action.

set -eo pipefail

set +e
source "$PWD/plz-out/gen/third_party/sh/shflags"
set -e

DEFINE_boolean dry_run "$FLAGS_FALSE" "Echo the Docker commands that would be run, rather than actually running them"

FLAGS "$@" || exit 1
eval set -- "$FLAGS_ARGV"

if [ "$FLAGS_dry_run" -eq "$FLAGS_TRUE" ]; then
  docker="echo docker"
else
  docker=docker
fi

default_repo="$(cat plz-out/gen/scripts/default_docker_repo)"

if [ -n "$GITHUB_REPOSITORY" ] && [ -n "$GITHUB_ACTOR" ] && [ -n "$GITHUB_TOKEN" ]; then
  ghcr_root="ghcr.io/$GITHUB_REPOSITORY"
fi

if [ -n "${DOCKERHUB_USERNAME}" ]; then
  echo "Authorising with Docker Hub"
  echo "${DOCKERHUB_PASSWORD}" | $docker login --username "${DOCKERHUB_USERNAME}" --password-stdin
fi

if [ -n "$ghcr_root" ]; then
  echo "Authorising with GitHub Container Registry"
  echo "$GITHUB_TOKEN" | $docker login ghcr.io --username "$GITHUB_ACTOR" --password-stdin
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
  $docker tag "${fqn}" "${fqn_version}"
  fqns_to_push+=("${fqn_version}")

  if [ -n "$ghcr_root" ]; then
    fqn="${fqn/$default_repo/$ghcr_root}"
    repo="$(echo "${fqn}" | cut -f1 -d\:)"
    fqn_version="${repo}:${version}"
    echo "-> tagging as ${fqn_version}"
    $docker tag "${fqn}" "${fqn_version}"
    fqns_to_push+=("${fqn_version}")
  fi
done

for fqn_to_push in "${fqns_to_push[@]}"; do
  echo "-> pushing as ${fqn_to_push}"
  $docker push "${fqn_to_push}"
done
