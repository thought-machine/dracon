subinclude("//build/defs:kustomize")

kustomized_config(
    name = "dev",
    srcs = [
        "elasticsearch-consumer.yaml",
        "enricher.yaml",
        "git-source.yaml",
        "gosec-producer.yaml",
        "kustomization.yaml",
        "pipeline.yaml",
        "pipeline-run.yaml",
    ],
    images = [
        "//consumers/elasticsearch:dracon-consumer-elasticsearch",
        "//cmd/enricher:dracon-enricher",
        "//source/git:dracon-source-git",
        "//producers/golang_gosec:dracon-producer-gosec",
    ],
)
