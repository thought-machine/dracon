sh_cmd(
    name = "go",
    srcs = ["go.sh"],
    cmd = """
source $(out_location //scripts/util)
export GO_LINT=$(out_location //third_party/lang:golint)
export GO_VET=$(out_location //third_party/lang:go_vet)
source $SRCS
    """,
    data = [
        "//scripts/util",
        "//third_party/lang:golint",
        "//third_party/lang:go_vet",
    ],
    labels = ["lint"],  # disable in `plz lint` until all are passing
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
    labels = ["lint"],
    shell = "/bin/bash",
)
