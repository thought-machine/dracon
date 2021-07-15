subinclude("//build/defs:kustomize")

kustomized_config(
    name = "base",
    srcs = [
        "kustomization.yaml",
        "tekton.dev_v1beta1_Pipeline_schema.json",
        ":openapi_schema",
        "test_schema.json",
        "pipeline.yaml",
        "transformer.yaml"
    ],
    visibility = ["//docs/designs/dracon-v2-example/..."],
)

sh_binary(
    name = "generate_openapi_schema",
    deps = ["//third_party/tools:jq"],
    main = "generate_openapi_schema.sh",
)

genrule(
    name = "openapi_schema",
    srcs = ["//third_party/k8s:tektoncd_v1beta1_swagger"],
    outs = ["openapi_schema.json"],
    tools = [":generate_openapi_schema"],
    cmd = "$TOOL",
)
