go_library(
    name = "gojira",
    srcs = [
        "gojira.go",
        "jirautils.go",
    ],
    deps = [
        "//third_party/go:go-jira",
        "//third_party/go:tgo",
        "//third_party/go:yaml.v2",
        "//consumers/jira_c:config",
    ],
    visibility = [
        "//consumers/jira_c/...",
    ],
)


go_test(
    name = "gojira_test",
    srcs = [
        "jirautils_test.go",
        "gojira_test.go",
    ],
    visibility = ["//consumers/jira_c/..."],
    deps = [
        ":gojira",
        "//third_party/go:go-jira",
        "//third_party/go:stretchr_testify",
    ],
)
