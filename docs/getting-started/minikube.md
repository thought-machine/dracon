## Getting Started with Dracon on Minikube

## Installation

A helper script that automates the below exists in `./scripts/minikube.sh`.

1. First install the latest release of Tekton Pipelines:

   ```bash
   $ kubectl --context minikube apply --filename https://storage.googleapis.com/tekton-releases/pipeline/previous/v0.9.0/release.yaml
   ```

2. Create a namespace for Dracon

   ```bash
   $ kubectl --context minikube create namespace dracon
   ```

3. Create a DB for the Enricher

   ```bash
   $ kubectl apply --context minikube --namespace dracon -f resources/persistence/enricher-db/k8s.yaml
   ```

   **Note**: Running postgres like this is not recommended in production, however it's great for a demo run of Dracon. You should use a properly set up and maintained Postgres instance with a secret username and password when you run Dracon in production.

4. Start Minio for ephemeral build storage

   ```bash
   $ kubectl apply --context minikube --namespace dracon -f resources/persistence/minio-storage/k8s.yaml
   ```

5. Create Elasticsearch+Kibana for the Elasticsearch Consumer to push results into

   ```bash
   $ kubectl apply --context minikube --namespace dracon -f resources/persistence/elasticsearch-kibana/elasticsearch.yaml
   $ kubectl apply --context minikube --namespace dracon -f resources/persistence/elasticsearch-kibana/kibana.yaml
   ```

6. Dracon is now ready to use. Check out the [Running Demos Guide](/docs/getting-started/tutorials/running-demos.md)

## Usage

### Configure Kubectl

Configure Kubectl to use the `minikube` context and `dracon` namespace by default:

```bash
$ kubectl config use-context minikube
$ kubectl config set-context minikube --namespace=dracon
```

### Setting up a Pipeline

To setup an pipeline, you can execute:

```bash
$ dracon setup --pipeline examples/pipelines/golang-project
```

### Running a Pipeline

To run that example pipeline you can execute:

```bash
$ dracon run --pipeline examples/pipelines/golang-project
```

### Inspecting a Pipeline

To see the progress of a Pipeline, you can execute:

```bash
# To see the pipeline status
$ kubectl get piplineruns
# To see the pods running as part of the pipeline
$ kubectl get pod -w
```
