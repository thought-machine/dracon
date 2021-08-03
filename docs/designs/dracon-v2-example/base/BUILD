subinclude("//build/defs:kustomize")

kustomized_config(
    name = "base",
    srcs = [
        "kustomization.yaml",
        "openapi_schema.json",
        "pipeline.yaml",
        "transformer.yaml",
    ],
    visibility = ["//docs/designs/dracon-v2-example/..."],
)

sh_binary(
    name = "generate_openapi_schema_tool",
    deps = ["//third_party/tools:jq"],
    main = ".generate_openapi_schema_tool.sh",
    visibility = ["//docs/..."]
)

genrule(
    name = "openapi_schema",
    srcs = ["//third_party/k8s:tektoncd_v1beta1_swagger"],
    outs = ["generated_openapi_schema.json"],
    tools = [":generate_openapi_schema_tool"],
    cmd = "$TOOL",
)

gentest(
    name = "openapi_schema_test",
    data = [
        "openapi_schema.json",
        ":openapi_schema",
    ],
    test_tools = ["//third_party/tools:jd"],
    test_cmd = """
output=$($TOOLS $(location openapi_schema.json) $(location :openapi_schema))
if [[ $output ]]; then
    echo "$output"
    exit 1
fi
    """,
    no_test_output = True,
    exit_on_error = True,
)
