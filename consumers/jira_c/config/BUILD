go_library(
    name = "config",
    srcs = [
        "config.go",
    ],
    deps = [
        ":types",
        "//third_party/go:yaml.v2",
    ],
    visibility = [
        "//consumers/jira_c:jira_c"
    ],
)

go_library(
    name = "types",
    srcs = [
        "types.go",
    ],
    visibility = [
        "//consumers/jira_c/..."
    ],
)

go_test(
    name = "config_test",
    srcs = [
        "config_test.go",
    ],
    deps = [
        ":config",
        ":types",
        "//third_party/go:stretchr_testify",
        "//third_party/go:yaml.v2",
    ],
)

filegroup(
    name = "config_file",
    srcs = ["config.yaml"],
    visibility = [
        "//consumers/jira_c:image",
    ]
)