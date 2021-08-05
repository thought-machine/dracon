# Dracon V2 :train::train:

Dracon's original interface has been around in its current state since the first release 0.1.0 ~1.5 years ago. Since then, we've become a bit wiser and more supporting tools have matured. This design document explores the improvements (some much needed) that we can make to help make Dracon more accessible and easy to use.

## Summary

- Upgrade and support the latest Tekton Pipelines using `tekton.dev/v1beta1`
  - `tekton.dev/v1alpha1.PipelineResource` is deprecated in favour of [_Tasks_][1].
  - Replace [Minio][2] with new [_Workspace_][3] features. TODO: use NFS, deploy NFS server alongside Tekton.
  - Use [Directed Acyclic Graph (DAG)][4] feature to order [_Tasks_][1].
- Use [Kustomize][5] instead of our own `dracon` patching and templating binary.
  - [_Components_][6] can be used to package Dracon _Producers_ and _Consumers_ as building blocks.
  - [_Pipelines_][7] can be composed via composition of [_Components_][6].
  - [_Components_][6] can be extended and modified via [_Patches_][8].
  - Kustomize [_Remote Build_][9] can be used to distribute [_Pipelines_][7].

[1]: https://github.com/tektoncd/pipeline/blob/v0.26.0/docs/tasks.md
[2]: https://min.io
[3]: https://github.com/tektoncd/pipeline/blob/v0.26.0/docs/workspaces.md
[4]: https://github.com/tektoncd/pipeline/blob/v0.26.0/docs/pipelines.md#configuring-the-task-execution-order
[5]: https://kustomize.io/
[6]: https://github.com/kubernetes-sigs/kustomize/blob/kustomize/v4.2.0/examples/components.md
[7]: https://github.com/tektoncd/pipeline/blob/v0.26.0/docs/pipelines.md
[8]: https://github.com/kubernetes-sigs/kustomize/blob/kustomize/v4.2.0/examples/patchMultipleObjects.md
[9]: https://github.com/kubernetes-sigs/kustomize/blob/master/examples/remoteBuild.md#url-format


### Tekton Pipelines with `tekton.dev/v1beta1`


#### Deprecation of `tekton.dev/v1alpha1.PipelineResource` 

#### Replacement of Minio with Tekton Workspace

#### Using Directed Acyclic Graph (DAG) to order Tekton Tasks


### Kustomize

#### Kustomize Components

##### Tekton Pipeline Composition with Kustomize Components

##### Extending and Modifying Kustomize Components via Kustomize Patches

#### Kustomize Remote Build

- GitHub release with base and components as artefacts.


## Improvements to Extending Dracon (Producers, Consumers and Enrichers)

* Please `dracon.build_defs` to generate Kustomize component.
  * Enforces pipeline interface expectations e.g.:
    * Adds pipeline patches for Producers that make them run before the Enricher(s).
  * Component gets published as an GitHub Actions artefact on build.
  * Component gets published as a GitHub Release artefact on release.

## Full Example


---
---
---
Below this is old/needs merging into above doc
---
---
---
---

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
