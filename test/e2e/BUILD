sh_cmd(
    name = "wait_for_pipelineruns",
    srcs = ["wait_for_pipelineruns.sh"],
    cmd = """
source "$(out_location //scripts/util)"
source $SRCS
    """,
    data = [
        "//scripts/util",
    ],
    shell = "/bin/bash",
)
