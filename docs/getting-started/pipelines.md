## Pipelines

Dracon utilises [Tekton Pipelines](https://tekton.dev/) to execute a given Dracon Pipeline.

There are example pipelines in the `./examples/pipelines` directory which are advised to be used as a base for your own Pipelines.

A Pipeline directory consists of all resources required to run a Pipeline which includes:

- Tekton Tasks
  - Dracon Producers
  - Dracon Consumers
  - Dracon Enrichers
  - Dracon Sources
- A Tekton Pipeline
- A Tekton PipelineRun
- Any further Tekton PipelineResources external to Dracon resources e.g. Git Source

### Using Secrets

Some tools, or even cloning your repository may require secrets. The following sections describe how to do this.

#### Git+SSH Cloning

An example `pipeline-run.yaml` for cloning via Git+SSH using Kubernetes Secrets is available at `./examples/git-ssh.pipeline-run.yaml` which provides the following resources:

- `Secret`: To store the "secret" data such as SSH private key and trusted host public key
- `ServiceAccount`: For Tasks within the Pipeline to assume this account and gain access to the secret
- `PipelineResource`: Utilise Tekton's built in git resource kind to fetch git source code
- `PipelineRun`: that uses the service account and git resource.

#### Using Secrets in Tasks

An example of changed resources to support using secrets in Tasks is available in `./examples/use-secrets.task.yaml` which provides the following resources:

- `Secret`: To store the "secret" data such as SSH private key and trusted host public key
- `ServiceAccount`: For Tasks within the Pipeline to assume this account and gain access to the secret
- `PipelineResource`: Utilise Tekton's built in git resource kind to fetch git source code
- `PipelineRun`: that uses the service account and git resource.
- `Task`: to use the secret in an environment variable.
