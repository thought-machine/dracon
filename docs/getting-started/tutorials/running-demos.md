# Dracon Demos

Example of running the the [Demo Dracon Pipelines](examples/pipelines/golang-project) where you get to see the results in kibana at the end.

## Prerequisites

- You have followed the [Getting Started with Minikube Guide](/docs/getting-started/minikube.md)

---

## Tutorial

1. You can run a demo pipeline with:

   ```bash
   $ dracon setup --pipeline examples/pipelines/golang-project
   $ dracon run --pipeline examples/pipelines/golang-project
   ```

2. Wait for the pipeline to finish running:

   ```bash
   $ kubectl get pipelineruns --watch
   ```

   Note: Use ctrl-c to exit the blocking watch

3. Once the pipelinerun has finished running you can view your results in Kibana (backed by Elasticsearch):

   ```bash
   $ kubectl port-forward svc/kibana 5601
   ```

   Visit http://localhost:5601, and create the dracon index filter
