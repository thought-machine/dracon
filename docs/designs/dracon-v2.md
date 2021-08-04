# Dracon V2

## Goals

- Latest Tekton w/ Tekton Dashboard
  - New Workspace features. TODO: use NFS, deploy NFS server alongside Tekton.
  - New Results features. Done.
  - PipelineResources are now replaced by Tasks.
- Kustomize.
  - Tasks:
    - Prepend containers. Done, works with Components too.
    - Add Volumes. Done, works with Components too.
- Kustomize Distribution. We can use https://github.com/kubernetes-sigs/kustomize/blob/master/examples/remoteBuild.md#url-format


## Kustomize

Currently, the Dracon binary is essentially just a glorified json-patch tool with some hardcoded features. We should try to replace this with Kustomize.

Given a `dracon-base` Kustomization, we can add overlays to modify the Kubernetes resources (Pipeline, Task, PipelineRun, etc.).

### Kustomize Components
See: https://github.com/kubernetes-sigs/kustomize/blob/master/examples/components.md

We propose to use Kustomize Components to package Dracon components (Producers, Consumers) which can be added to a Dracon pipeline via composition.

For example, a file tree may look like:

```
./components
./pipelines
./pipelines/base
./pipelines/golang-example
./pipelines/mixed-example
./pipelines/python-example
```

This allows us to focus on creating re-usable components that end-users can use to easily build their own pipelines.

### Modification of `tekton.dev/v1beta1/Pipeline`

https://github.com/tektoncd/pipeline/blob/main/docs/pipelines.md#passing-one-tasks-results-into-the-parameters-or-whenexpressions-of-another

#### Examples
##### Adding a Dracon Producer

```yaml
---
# ./components/my-producer/patch-pipeline.yaml
apiVersion: tekton.dev/v1beta1
kind: Pipeline
metadata:
  name: pipeline-with-parameters
spec:
  tasks:
  - name: my-producer
    taskRef:
      name: my-producer
  # add output of task as parameter to the enricher to make the enricher dependent on this.
  - name: enricher
    params: 
    - name: producers
      value: 
      - "$(tasks.my-producer.results.issues)"
---
# ./components/my-producer/my-producer.yaml
apiVersion: tekton.dev/v1beta1
kind: Task
metadata:
  name: my-producer
spec:
  results:
  - name: issues
    description: The issues found by this Dracon producer.
  steps:
  - name: produce-issues
    image: bash:latest
    script: |
      #!/usr/bin/env bash
      date +%s | tee $(results.issues.path)
---
# ./components/my-producer/kustomization.yaml
---
apiVersion: kustomize.config.k8s.io/v1alpha1
kind: Component

resources:
  - my-producer.yaml

patches:
  - path: patch-pipeline.yaml
    target:
      kind: Pipeline
```

##### Adding a Dracon Consumer

```yaml
# patches/add-my-consumer.yaml
---
apiVersion: tekton.dev/v1beta1
kind: Pipeline
metadata:
  name: pipeline-with-parameters
spec:
  tasks:
  - name: my-consumer
    taskRef:
      name: my-consumer
    params:
    # use issues output from enricher to create dependency on enricher.
    - name: enriched-issues
      value: "$(tasks.enricher.results.issues)"
```


TODO: How can we make this easier?
 - Define an interface for people to implement?
   - Producer Tasks must output `issues` as a `result`.
   - Producer Tasks must be defined in the Pipeline with their `issues` result as a parameter to the `enricher`.
   - **Optional**: A Please build def which generates the yaml? just to keep code DRY.

## Kustomize Distribution

End users should be able to use Kustomize to modify any of our example pipelines for their own use-cases. e.g.

```yaml
# ./kustomization.yaml
resources:
- github.com/thought-machine/dracon/examples/pipelines/golang-project?ref=v0.12.1
- my-producer.yaml
# if the end-user is using plz, they can use a remote_file and reference the output directory instead

namespace: dracon-demo
commonLabels:
  my-org.io/team: my-team

patches:
- path: patches/add-my-producer.yaml
  target:
    group: tekton.dev
    version: v1beta1
    kind: Pipeline
```
