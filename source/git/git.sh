#!/bin/sh -xe

git_src="/workspace/git-source"
out_src="/workspace/output/source/source.tgz"

cd "${git_src}"

addr=$(git remote -v | cut -f1 -d" " | cut -f2 | head -n1 | cut -f2 -d"@")
rev=$(git rev-parse HEAD)

echo "${addr}?ref=${rev}" > .source.dracon

tar -C "${git_src}/" -czf "${out_src}" .
