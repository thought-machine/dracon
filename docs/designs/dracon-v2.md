# Dracon V2

## Goals

- Latest Tekton w/ Tekton Dashboard
  - New Workspace features.
  - New Results features.
  - PipelineResources are now replaced by Tasks.
- Kustomize.
- Kustomize Distribution.


## Kustomize

Currently, the Dracon binary is essentially just a glorified json-patch tool with some hardcoded features. We should try to replace this with Kustomize.

Given a `dracon-base` Kustomization, we can add overlays to modify the Kubernetes resources (Pipeline, Task, PipelineRun, etc.).


### Modification of `tekton.dev/v1beta1/Pipeline`

https://github.com/tektoncd/pipeline/blob/main/docs/pipelines.md#passing-one-tasks-results-into-the-parameters-or-whenexpressions-of-another

#### Examples
##### Adding a Dracon Producer

```yaml
---
# patches/add-my-producer.yaml
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
# resources/my-producer.yaml
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
```

TODO: How can we make this easier?
 - Define an interface for people to implement?
   - Producer Tasks must output `issues` as a `result`.
   - Producer Tasks must be defined in the Pipeline with their `issues` result as a parameter to the `enricher`.
 - A [Kustomize plugin](https://kubectl.docs.kubernetes.io/guides/extending_kustomize/)?
   - create `dracon.thoughtmachine.net/v1/Pipeline`

```yaml
---
apiVersion: dracon.thoughtmachine.net/v1
kind: Pipeline
metadata:
  name: notImportantHere
producers:
  - my-producer
consumers:
  - my-consumer
# This would generate a tekton.dev/v1beta1/Pipeline with the given producers and consumers.
# It would also be able to provide build-time validation (hopefully helpful output).
# Therefore end-users would only need to modify the Pipeline if modifying the enricher or retrieval of sourcecode.

# TODO: does this work with overlays? i.e. can we generate a Pipeline then use overlay kustomizations on it (e.g. change the task which retrieves source-code)?
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

## Kustomize Distribution

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
- path: patches/my-producer.yaml
  target:
    group: tekton.dev
    version: v1beta1
    kind: Task
```
