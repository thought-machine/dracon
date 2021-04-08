sh_cmd(
    name = "go",
    srcs = ["go.sh"],
    cmd = """
source $(out_location //scripts/util)
export GO_ROOT=$(out_location //third_party/lang:go_tool)
export GO_FMT="\\\\$GO_ROOT/bin/gofmt"
source $SRCS
    """,
    data = [
        "//scripts/util",
        "//third_party/lang:go_tool",
    ],
    labels = ["fmt"],
    shell = "/bin/bash",
)

sh_cmd(
    name = "plz",
    srcs = ["plz.sh"],
    cmd = """
source $(out_location //scripts/util)
source $SRCS
    """,
    data = [
        "//scripts/util",
    ],
    labels = ["fmt"],
    shell = "/bin/bash",
)
