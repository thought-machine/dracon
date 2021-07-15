subinclude("//build/defs:kustomize")

kustomized_config(
    name = "golang-example",
    srcs = [
        "kustomization.yaml",
        "gosec-producer.yaml",
        "second-producer.yaml",
        "patches/add-gosec-to-pipeline.yaml",
        "patches/add-second-to-pipeline.yaml",
    ],
    deps = [
        "//docs/designs/dracon-v2-example/dracon-base/base",
    ],
)
