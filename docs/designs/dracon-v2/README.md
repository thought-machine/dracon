# Dracon V2 :train::train:

Dracon's original interface has been around in its current state since the first release 0.1.0 ~1.5 years ago. Since then, we've become a bit wiser and more supporting tools have matured. This design document explores the improvements (some much needed) that we can make to help make Dracon more accessible and easy to use.

## Summary

- Upgrade and support the latest version of [Tekton Pipelines][0] using `tekton.dev/v1beta1`
  - `tekton.dev/v1alpha1.PipelineResource` is deprecated in favour of [_Tasks_][1].
  - Replace [Minio][2] with new [_Workspace_][3] features. TODO: use ceph, deploy ceph with [Rook](https://rook.io/)
  - Use [Directed Acyclic Graph (DAG)][4] feature to order [_Tasks_][1].
- Use [Kustomize][5] instead of our own `dracon` patching and templating binary.
  - [_Components_][6] can be used to package Dracon _Producers_ and _Consumers_ as building blocks.
  - [_Pipelines_][7] can be composed via composition of [_Components_][6].
  - [_Components_][6] can be extended and modified via [_Patches_][8].
  - Kustomize [_Remote Build_][9] can be used to distribute [_Pipelines_][7].

[0]: https://github.com/tektoncd/pipeline
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

We should upgrade our use of Tekton Pipelines to the newest API Version `tekton.dev/v1beta1` to take advantage of the newer features.

#### Deprecation of `tekton.dev/v1alpha1.PipelineResource` 

`tekton.dev/v1alpha1.PipelineResource`  did not graduate into `tekton.dev/v1beta1` and they have been replaced with `tekton.dev/v1beta1.Task`s.

#### Replacement of Minio with Tekton Workspace

We are currently using Minio to support our use of `tekton.dev/v1alpha1.PipelineResource`s. We can replace this with Tekton Workspaces which uses Kubernetes Volumes to share filesystems between Tasks in a Pipeline.
TODO: play with NFS

#### Using Directed Acyclic Graph (DAG) to order Tekton Tasks

We currently use `tekton.dev/v1alpha1.Pipeline.spec.tasks[].runAfter` to order tasks within a Pipeline. This makes composing `Pipeline`s from reusable parts very difficult. For Dracon V2, we should use `"$(tasks.my-task.results.issues)"` in `tekton.dev/v1beta1.Pipeline.spec.tasks[].params` to create a Directed Acyclic Graph (DAG) based on Task's parameters and results. 

For example:

```yaml
---
# task-a.yaml
apiVersion: tekton.dev/v1beta1
kind: Task
metadata:
  name: task-a
spec:
  results:
  - name: result-a
  steps: []
---
# task-b.yaml
apiVersion: tekton.dev/v1beta1
kind: Task
metadata:
  name: task-b
spec:
  params:
  - name: param-a
    type: array
  results:
  - name: result-a
  steps: []
---
# my-pipeline.yaml
apiVersion: tekton.dev/v1beta1
kind: Pipeline
metadata:
  name: my-pipeline
spec:
  tasks:
  - name: task-a
    taskRef: task-a
  - name: task-b
    taskRef: task-b
    params:
    - name: param-a
      value: 
      - "$(tasks.task-a.results.result-a)"
```

will schedule `task-a -> task-b` as `task-b` depends on the results of `task-a`.

### Kustomize

[Kustomize](https://kustomize.io/) is a tool to perform YAML patching of Kubernetes resources. The current `dracon` CLI tool has some similarities but much fewer features. For Dracon V2, we should replace the `dracon` CLI tool with Kustomize.

#### Kustomize Components

Kustomize Components is a way for us to create re-usable "components" that you can add to Kustomize configurations.
##### Tekton Pipeline Composition with Kustomize Components

In Dracon, we construct Tekton Pipelines that consist of Dracon components: Producers, Consumers and Enrrichers. We can create a Kustomize component for each Dracon Producer, Consumer and Enricher which can then be used to compose Dracon Pipelines (which use Tekton Pipelines).

For example:

```yaml
---
# ./components/my-producer/kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1alpha1
kind: Component

resources:
  - my-producer.yaml

patches:
  - path: patch-pipeline.yaml
    target:
      kind: Pipeline
---
# ./components/my-producer/my-producer.yaml
apiVersion: tekton.dev/v1beta1
kind: Task
metadata:
  name: my-producer
spec:
  results:
  - name: issues
    description: The path for the issues found by this Dracon producer.
  steps:    
  - name: produce-issues
    image: bash:latest
    script: echo "$HOME/producer.out" > $(results.issues.path)
---
# ./components/my-producer/patch-pipeline.yaml
apiVersion: tekton.dev/v1beta1
kind: Pipeline
metadata:
  name: unused
spec:
  tasks:
  # add my-producer Task to the Pipeline.
  - name: my-producer
    taskRef:
      name: my-producer
    params: 
    - name: source-path
      value: "$(tasks.get-source.results.url)"
  # add output of my-producer as parameter to the enricher to make the enricher depend on my-producer.
  - name: enricher
    params: 
    - name: producer-result-paths
      value: 
      - "$(tasks.my-producer.results.issues)"
```

We can template the `patch-pipeline.yaml` per Producer/Consumer/Enricher in a `dracon.build_def` as these patches will be very similar per component.

##### Extending and Modifying Kustomize Components via Kustomize Patches

Kustomize Components can also be affected by Kustomize Patches which we can leverage to enable end-users to customize pre-defined Producers/Consumers/Enrichers in their Pipelines.

#### Kustomize Remote Build

Kustomize Remote Build enables end-users to directly reference remote Kustomize configuration. 

```yaml
# ./kustomization.yaml
resources:
- github.com/thought-machine/dracon/pipelines/base?ref=v2.0.0

components:
- github.com/thought-machine/dracon/components/golang-producer?ref=v2.0.0
- github.com/thought-machine/dracon/components/enricher?ref=v2.0.0
- github.com/thought-machine/dracon/components/elastichsearch-consumer?ref=v2.0.0

namespace: dracon-demo
commonLabels:
  my-org.io/team: my-team
```

**Note**: The above isn't tested but will be along the lines of what we're aiming for.


## Improvements to Extending Dracon (Producers, Consumers and Enrichers)

* Please `dracon.build_defs` to generate Kustomize component.
  * Enforces pipeline interface expectations e.g.:
    * Adds pipeline patches for Producers that make them run before the Enricher(s).
    * Adds pipeline patches for Consumers that make them run after the Enricher(s).
    * Adds pipeline patches for Enrichers that make them run after the Producer(s) and before the Consumer(s).
  * Component gets published as an GitHub Actions artefact on build.
  * Component gets published as a GitHub Release artefact on release.

The published artefacts can then be referenced via Kustomize Remote Build.

## Full Guided Example

1. Run the helper script which creates and configures a KinD cluster.

    ```bash
    $ ./pleasew run //docs/designs/dracon-v2/scripts/kind:setup
    ```

2. Create a namespace for Dracon

   ```bash
   $ ./pleasew deploy //scripts/development/k8s:namespace
   ```

3. Create DB for the Enricher

   ```bash
   $ ./pleasew deploy //scripts/development/k8s/enricher-db --wait
   ```
  
4. Setup example Dracon V2 pipeline
  
    ```bash
    $ ./pleasew run //docs/designs/dracon-v2/examples/golang-project:golang-project_setup
    ```

5. Run example Dracon V2 pipeline
  
    ```bash
    $ ./pleasew run //docs/designs/dracon-v2/examples/golang-project:golang-project_run
    ```

6. Observe Pipeline progress in browser at http://tekton-dashboard.localhost
