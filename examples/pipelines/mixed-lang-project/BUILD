subinclude("//build/defs:kustomize")

kustomized_config(
    name = "dev",
    srcs = [
        "elasticsearch-consumer.yaml",
        "enricher.yaml",
        "git-source.yaml",
        "gosec-producer.yaml",
        "spotbugs-producer.yaml",
        "bandit-producer.yaml",
        "kustomization.yaml",
        "pipeline-run.yaml",
        "pipeline.yaml",
    ],
    images = [
        "//consumers/elasticsearch_c:dracon-consumer-elasticsearch",
        "//cmd/enricher:dracon-enricher",
        "//source/git:dracon-source-git",
        "//producers/golang_gosec:dracon-producer-gosec",
        "//producers/python_bandit:dracon-producer-bandit",
        "//producers/java_spotbugs:dracon-producer-java_spotbugs",
    ],
)
