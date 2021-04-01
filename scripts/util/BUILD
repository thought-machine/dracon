sh_cmd(
    name = "util",
    srcs = ["util.sh"],
    cmd = """
source "$(out_location //third_party/sh:ansi)"
YQ_BIN="$(out_location //third_party/tools:yq)"
source $SRCS
    """,
    data = [
        "//third_party/sh:ansi",
        "//third_party/tools:yq",
    ],
    visibility = [
        "//scripts/...",
        "//test/...",
    ],
)
