go_library(
    name = "config",
    srcs = [
        "config.go",
        "types.go",
    ],
    deps = [
        "//third_party/go:yaml.v2",
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
        "//third_party/go:stretchr_testify",
        "//third_party/go:yaml.v2",
    ],
)
