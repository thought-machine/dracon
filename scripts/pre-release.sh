#!/bin/bash
set -euo pipefail

version=$(git describe --always)

github_api_version="application/vnd.github.v3+json"
github_owner="thought-machine"
github_repo="dracon"

post_release_json=$(cat <<EOF
{
  "tag_name": "${version}",
  "target_commitish": "master",
  "name": "${version}",
  "body": "",
  "draft": false,
  "prerelease": true
}
EOF
)

release_resp=$(curl \
  --header "Accept: ${github_api_version}" \
  --header "Authorization: token ${GITHUB_TOKEN}" \
  --silent \
  --request POST \
  --data "${post_release_json}" \
  "https://api.github.com/repos/${github_owner}/${github_repo}/releases")

release_id=$(echo "${release_resp}" | jq '.id')
if [ -z "${release_id}" ] || [ "${release_id}" == "null" ]; then
  echo "!> could not find release id in response"
  echo "${release_resp}"
  exit 1
fi
echo "-> created release ${release_id}"

assets=$(find plz-out/bin/cmd/dracon -executable -type f \( ! -iname ".*" ! -iname "*_test" \))
for asset in ${assets}; do
  asset_name="${asset//plz-out\/bin\/cmd\//}"
  asset_name=$(dirname "${asset_name}")
  asset_name="${asset_name//\//_}"
  asset_resp=$(curl \
  --header "Accept: ${github_api_version}" \
  --header "Authorization: token ${GITHUB_TOKEN}" \
  --header "Content-Type: application/octet-stream" \
  --silent \
  --request POST \
  --data-binary @"${asset}" \
  "https://uploads.github.com/repos/${github_owner}/${github_repo}/releases/${release_id}/assets?name=${asset_name}")
  asset_id=$(echo "${asset_resp}" | jq '.id')
  if [ -z "${asset_id}" ] || [ "${asset_id}" == "null" ]; then
    echo "!> could not find asset id in response"
    echo "${asset_resp}"
    exit 1
  fi
  echo "uploaded ${asset_name} to release ${release_id}"
done

echo ""
echo ""
echo "-> Created pre-release ${version} and uploaded assets"
