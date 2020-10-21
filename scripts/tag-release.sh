#!/bin/bash
set -euo pipefail

git fetch origin master --tags &> /dev/null
commit_sha=$(git rev-parse origin/master)

highest_tag=$(git tag | sort | tail -n1)

major=$(echo "${highest_tag}" | cut -f1 -d. | tr -dc '0-9')
minor=$(echo "${highest_tag}" | cut -f2 -d. | tr -dc '0-9')
patch=$(echo "${highest_tag}" | cut -f3 -d. | tr -dc '0-9')

printf "> The highest current tag is '%s' (major: %d, minor: %d, patch: %d)\n" "${highest_tag}" "${major}" "${minor}" "${patch}"
printf "> What kind of release is this? (For guidance, see: https://semver.org/) [major/minor/patch] " 
read release_type
case "${release_type}" in
major)
    major=$((major+1))
    minor=0
    patch=0
    ;;
minor)
    minor=$((minor+1))
    patch=0
    ;;
patch)
    patch=$((patch+1))
    ;;
*)
  printf "!> Invalid option: '%s'.\n" "${release_type}"
  exit 1
  ;;
esac

new_tag="v${major}.${minor}.${patch}"

printf "> The new release will be '%s'. Is this OK? [y/N] " "${new_tag}"
read ok
if [ "${ok}" != "y" ]; then
    printf "Not OK, exiting.\n"
fi

git tag --annotate "${new_tag}" --message "${new_tag}" "${commit_sha}"
git push origin --tags
