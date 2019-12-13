# Releases

Dracon releases follow [Semantic Versioning](https://semver.org/)

We use annotated git tags to define a release by doing the following:

```
$ git tag --annotate <version>
e.g:
# git tag --annotate 0.1.0
# then push the tags
$ git push --tags
```

This lets us use `git desribe` to give us a descriptive version from any commit in the format `<version>-<commits since version>-<short-sha>`, e.g.:

- `0.0.0` (when on the annotated commit (`git checkout 0.0.0`))
- `0.0.0-3-g7487887`

## Creating a Release

1. Tag the commit you would like to release from `master` and push the tags:
   ```
   git fetch --tags
   git tag --annotate <version> --message "<version>"
   git push origin --tags
   ```
2. Build and publish the release Docker images to Docker Hub by running:
   ```
   # TODO(vj): make this script push image tags MAJOR and MINOR too. e.g. pushing tag `:0.1.2` should also push tags `:0` and `:0.1`
   scripts/publish-images.sh <version>
   # if this should be considered the "latest" version, push as the latest tag too:
   scripts/publish-images.sh latest
   ```
3. Build the Dracon binary and publish it to GitHub releases by running:
   ```
   scripts/publish-pre-release.sh
   ```
4. In the GitHub web interface, you can then promote the new pre-release to a release
