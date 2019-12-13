# Running Dracon Demo Pipeline Against a Public Repository via Git+HTTPS

## Prerequisites

- You have followed the [Getting Started with Minikube Guide](/docs/getting-started/minikube.md)

---

## Tutorial

1. Clone the Dracon repository.

   ```bash
   $ git clone https://github.com/thought-machine/dracon-private.git "${PWD}/dracon"
   ```

2. Copy an example demo pipeline from the `./examples/pipelines` directory into your working directory. We have chosen `mixed-lang-project`.
   ```bash
   $ cp -r "${PWD}/dracon/examples/pipelines/mixed-lang-project" "${PWD}"
   ```
3. Update the `tekton.dev/v1alpha1, PipelineResource` in `pipeline-run.yaml`:

   1. Set `spec.params[0].value` to your desired git revision/branch.
   2. Set `spec.params[1].value` to your desired git public git url.

      ```yaml
      ---
      # git+https config
      apiVersion: dracon/v1alpha1
      kind: PipelineResource
      metadata:
        name: "{{.RunID}}-git-github-oauth2_proxy"
        labels: {}
      spec:
        type: git
        params:
          - name: revision
            value: master
          - name: url
            value: https://github.com/pusher/oauth2_proxy.git
      ```

4. You can now run `dracon setup` and `dracon run` with your pipeline

   ```bash
   $ dracon setup --pipeline mixed-lang-project
   $ dracon run --pipeline mixed-lang-project
   ```
