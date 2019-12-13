#!/bin/bash -e

version=$(git describe --always)

echo "-> Enter your GitHub API Personal Access Token"
github_api_token=""
read -r -p "" github_api_token
# github_api_token=$(echo "${github_api_token}" | tr -d '[:space:]')

echo "-> Creating pre-release version: ${version}"
echo "   Is this OK? (Ctrl+C to cancel, Enter to continue)"
read -p ""

plz build //cmd/dracon:dracon

github_api_version="application/vnd.github.v3+json"
github_owner="thought-machine"
github_repo="dracon-private"

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
  --header "Authorization: token ${github_api_token}" \
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

asset_resp=$(curl \
  --header "Accept: ${github_api_version}" \
  --header "Authorization: token ${github_api_token}" \
  --header "Content-Type: application/octet-stream" \
  --silent \
  --request POST \
  --data-binary @"${PWD}/plz-out/bin/cmd/dracon/dracon" \
  "https://uploads.github.com/repos/${github_owner}/${github_repo}/releases/${release_id}/assets?name=dracon")
asset_id=$(echo "${asset_resp}" | jq '.id')
if [ -z "${asset_id}" ] || [ "${asset_id}" == "null" ]; then
  echo "!> could not find asset id in response"
  echo "${asset_resp}"
  exit 1
fi

echo ""
echo ""
echo "-> Created pre-release ${version} and uploaded dracon assets"
