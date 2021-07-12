filegroup(
    name = "base",
    srcs = glob(["**.yaml", "*.json"]),
    visibility = ["PUBLIC"],
)

genrule(
    name = "base_k8s",
    srcs = [":base"],
    outs = ["base_k8s.yaml"],
    cmd = """
$TOOLS build $PKG > $OUTS
    """,
    tools = ["//docs/designs/dracon-v2-example/third_party/binary:kustomize"],
)
