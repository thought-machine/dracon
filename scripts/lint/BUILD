sh_cmd(
    name = "go",
    srcs = ["go.sh"],
    cmd = """
source $(out_location //scripts/util)
export GO_LINT=$(out_location //third_party/lang:golint)
source $SRCS
    """,
    data = [
        "//scripts/util",
        "//third_party/lang:golint",
    ],
    shell = "/bin/bash",
    # labels = ["lint"], disable in `plz lint` until all are passing
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
    labels = ["lint"],
    shell = "/bin/bash",
)
