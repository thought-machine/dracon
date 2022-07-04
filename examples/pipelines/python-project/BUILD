subinclude("//build/defs:kustomize")

kustomized_config(
    name = "dev",
    srcs = [
        "bandit-producer.yaml",
        "elasticsearch-consumer.yaml",
        "enricher.yaml",
        "git-source.yaml",
        "kustomization.yaml",
        "pipeline.yaml",
        "pipeline-run.yaml",
    ],
    images = [
        "//consumers/elasticsearch:dracon-consumer-elasticsearch",
        "//cmd/enricher:dracon-enricher",
        "//source/git:dracon-source-git",
        "//producers/python_bandit:dracon-producer-bandit",
    ],
)
