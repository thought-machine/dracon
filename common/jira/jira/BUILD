go_library(
    name = "jira",
    srcs = [
        "api.go",
        "apiutils.go",
    ],
    visibility = [
        "PUBLIC",
    ],
    deps = [
        "//common/jira/config",
        "//third_party/go:go-jira",
        "//third_party/go:tgo",
    ],
)

go_test(
    name = "jira_test",
    srcs = [
        "api_test.go",
        "apiutils_test.go",
    ],
    deps = [
        ":jira",
        "//third_party/go:go-jira",
        "//third_party/go:stretchr_testify",
    ],
)
