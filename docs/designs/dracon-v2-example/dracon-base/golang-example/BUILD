filegroup(
    name = "golang-example",
    srcs = glob(["**.yaml", "*.json"]),
    visibility = ["PUBLIC"],
)

genrule(
    name = "golang-example_k8s",
    srcs = [":golang-example", "//docs/designs/dracon-v2-example/dracon-base/base"],
    outs = ["golang-example_k8s.yaml"],
    cmd = """
$TOOLS build $PKG > $OUTS
    """,
    tools = ["//docs/designs/dracon-v2-example/third_party/binary:kustomize"],
)
