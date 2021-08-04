subinclude("//build/defs:kustomize")

kustomized_config(
    name = "golang-example",
    srcs = [
        "kustomization.yaml",
    ],
    deps = [
        "//docs/designs/dracon-v2-example/base",
        "//docs/designs/dracon-v2-example/components/gosec-producer",
    ],
)

sh_test(
    name = "golang-example_test",
    src = "test.sh",
    data = [
        ":golang-example_k8s",
        "//third_party/sh:shunit2",
        "//third_party/tools:yq",
    ],
)
