subinclude("//build/defs:kustomize")

kustomized_config(
    name = "golang-project",
    srcs = [
        "kustomization.yaml",
    ],
    deps = [
        "//docs/designs/dracon-v2/base",
        "//docs/designs/dracon-v2/components/gosec-producer",
    ],
)

sh_test(
    name = "golang-project_test",
    src = "test.sh",
    data = {
        "SHUNIT": ["//third_party/sh:shunit2"],
        "YQ": ["//third_party/tools:yq"],
        "INPUT": [":golang-project_k8s"],
    },
)
