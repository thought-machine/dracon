# Design: Kustomize Support and Motivation

[Kustomize](https://kustomize.io/) is a highly popular tool for mutating/patching Kubernetes configuration files.

## Local Development
We can patch development parameters in transparently via plz. For example, `imagePullPolicy: Never` so that local clusters only use images that have been preloaded into the cluster.

## WIP: Distribution

_Note_: this section is WIP and subject to changes.

End users should be able to use Kustomize to modify any of our example pipelines for their own use-cases. e.g.

```yaml
# ./kustomization.yaml
resources:
- github.com/thought-machine/dracon/examples/pipelines/golang-project?ref=v0.12.1
# if the end-user is using plz, they can use a remote_file and reference the output directory instead

namespace: dracon-demo
commonLabels:
    my-org.io/team: my-team

patches:
- path: patches/my-repository.yaml
  target:
    group: dracon
    version: v1beta1
    kind: PipelineResource
    name: "{{.RunID}}-git-github-oauth2-proxy"
# note: this above patch doesn't currently read well so we may need to reconsider how we define pipelines.

# ./patches/my-repository.yaml
---
apiVersion: dracon/v1beta1
kind: PipelineResource
metadata:
  name: "{{.RunID}}-git-github-oauth2-proxy"
  labels: {}
spec:
  type: git
  params:
  - name: revision
    value: master
  - name: url
    value: https://github.com/my-org/my-repository.git
```

running `kustomize build` will yield the kustomized configuration to stdout, which can be redirected to a file.

e.g.

```
$ mkdir -p modified-golang-project/
$ kustomize build > modified-golang-project/k8s_kustomized.yaml
$ dracon setup --pipeline modified-golang-project
$ dracon run --pipeline modified-golang-project
```
