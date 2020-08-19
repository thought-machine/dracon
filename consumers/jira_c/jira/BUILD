go_library(
    name = "jira",
    srcs = [
        "api.go",
        "apiutils.go",
    ],
    deps = [
        "//third_party/go:go-jira",
        "//third_party/go:tgo",
        "//consumers/jira_c/types:config",
    ],
    visibility = [
        "//consumers/jira_c/...",
    ],
)


go_test(
    name = "jira_test",
    srcs = [
        "api_test.go",
        "apiutils_test.go",
    ],
    visibility = ["//consumers/jira_c/..."],
    deps = [
        ":jira",
        "//third_party/go:go-jira",
        "//third_party/go:stretchr_testify",
    ],
)
