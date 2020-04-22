version = "3.6.0"

remote_file(
    name = "protoc_zip",
    out = "protoc-%s.zip" % version,
    url = "https://github.com/google/protobuf/releases/download/v%s/protoc-%s-$XOS-$XARCH.zip" % (version, version),
)

genrule(
    name = "protoc",
    srcs = [":protoc_zip"],
    outs = ["protoc"],
    binary = True,
    cmd = "$TOOL x $SRCS bin/protoc",
    tools = [CONFIG.JARCAT_TOOL],
    visibility = ["PUBLIC"],
    deps = [":protoc_deps"],
)

genrule(
    name = "protoc_deps",
    srcs = [":protoc_zip"],
    outs = ["include"],
    binary = True,
    cmd = "$TOOL x $SRCS",
    tools = [CONFIG.JARCAT_TOOL],
    visibility = ["PUBLIC"],
)
