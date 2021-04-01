## Getting Started with Dracon on KinD

## Installation

1. Run the helper script which creates and configures a KinD cluster.

    ```bash
    $ ./pleasew run //scripts/development/kind:setup
    ```

2. Create a namespace for Dracon

   ```bash
   $ ./pleasew deploy //scripts/development/k8s:namespace
   ```

3. Create a DB for the Enricher

   ```bash
   $ ./pleasew deploy //scripts/development/k8s/enricher-db --wait
   ```

   **Note**: Running postgres like this is not recommended in production, however it's great for a demo run of Dracon. You should use a properly set up and maintained Postgres instance with a secret username and password when you run Dracon in production.

4. Start Minio for ephemeral build storage

   ```bash
   $ ./pleasew deploy //scripts/development/k8s/minio-storage --wait
   ```

5. Create Elasticsearch+Kibana for the Elasticsearch Consumer to push results into

   ```bash
   $ ./pleasew deploy //scripts/development/k8s/elasticsearch-kibana --wait
   ```

6. Dracon is now ready to use. Check out the [Running Demos Guide](/docs/getting-started/tutorials/running-demos.md)

### Running a Pipeline

To run that example pipeline you can execute:

```bash
$ dracon run --context kind-dracon --namespace=dracon run --pipeline examples/pipelines/golang-project
```

_Note: make sure you have installed dracon via [Dracon - Installing](https://github.com/thought-machine/dracon#installing)_

### Inspecting a Pipeline

To see the progress of a Pipeline, you can execute:

```bash
# To see the pipeline status
$ kubectl -n dracon get pipelineruns
# To see the pods running as part of the pipeline
$ kubectl -n dracon get pod -w
```
