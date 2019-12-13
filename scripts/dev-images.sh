#!/bin/bash -e

plz query alltargets //... --include docker-build | sed 's/$/_load/g' | plz -p -v 2 --colour run sequential

# get all tags
plz query alltargets //... --include docker-build | sed 's/$/_fqn/g' | plz -p -v 2 --colour build
all_tag_files=$(find . -type f -name "*_fqn")
all_tags=""
for tag_file in ${all_tag_files}; do
  tag=$(cat ${tag_file})
  all_tags+=" ${tag}"
done

# Retag as latest
for tag in ${all_tags}; do
  repository=$(echo "${tag}" | cut -f1 -d":")
  new_tag="${repository}:latest"
  docker tag "${tag}" "${new_tag}"
  echo "tagged ${new_tag}"
done
