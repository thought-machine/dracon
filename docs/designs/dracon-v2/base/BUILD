subinclude("//build/defs:kustomize")

kustomized_config(
    name = "base",
    srcs = [
        "kustomization.yaml",
        "openapi_schema.json",
        "pipeline.yaml",
        "transformer.yaml",
    ],
    visibility = ["//docs/designs/dracon-v2/..."],
)

sh_binary(
    name = "generate_openapi_schema_tool",
    main = ".generate_openapi_schema_tool.sh",
    visibility = ["//docs/designs/dracon-v2/..."],
    deps = ["//docs/designs/dracon-v2/third_party/binary:jq"],
)

genrule(
    name = "openapi_schema",
    srcs = ["//docs/designs/dracon-v2/third_party/api:tektoncd_v1beta1_swagger"],
    outs = ["generated_openapi_schema.json"],
    cmd = "$TOOL",
    tools = [":generate_openapi_schema_tool"],
)

gentest(
    name = "openapi_schema_test",
    data = [
        "openapi_schema.json",
        ":openapi_schema",
    ],
    exit_on_error = True,
    no_test_output = True,
    test_cmd = """
output=$($TOOLS $(location openapi_schema.json) $(location :openapi_schema))
if [[ $output ]]; then
    echo "$output"
    exit 1
fi
    """,
    test_tools = ["//docs/designs/dracon-v2/third_party/binary:jd"],
)
