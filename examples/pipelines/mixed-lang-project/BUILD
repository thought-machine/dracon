subinclude("//build/defs:kustomize")

kustomized_config(
    name = "dev",
    srcs = [
        "bandit-producer.yaml",
        "elasticsearch-consumer.yaml",
        "enricher.yaml",
        "git-source.yaml",
        "gosec-producer.yaml",
        "kustomization.yaml",
        "pipeline.yaml",
        "pipeline-run.yaml",
        "spotbugs-producer.yaml",
    ],
    images = [
        "//consumers/elasticsearch:dracon-consumer-elasticsearch",
        "//cmd/enricher:dracon-enricher",
        "//source/git:dracon-source-git",
        "//producers/golang_gosec:dracon-producer-gosec",
        "//producers/python_bandit:dracon-producer-bandit",
        "//producers/java_spotbugs:dracon-producer-java_spotbugs",
    ],
)
