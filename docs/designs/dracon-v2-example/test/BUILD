subinclude("//build/defs:kustomize")

kustomized_config(
    name = "_test",
    srcs = [
        "kustomization.yaml",
    ],
    deps = [
        "//docs/designs/dracon-v2-example/base",
        "//docs/designs/dracon-v2-example/components/test-a",
        "//docs/designs/dracon-v2-example/components/test-b",
    ],
)


sh_test(
    name = "test",
    src = "test.sh",
    data = {
        "SHUNIT": ["//third_party/sh:shunit2"],
        "YQ": ["//third_party/tools:yq"],
        "INPUT": [":_test_k8s"],
    }
)
