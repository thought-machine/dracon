go_library(
    name = "jira",
    srcs = [
        "api.go",
        "apiutils.go",
    ],
    deps = [
        "//third_party/go:go-jira",
        "//third_party/go:tgo",
        "//common/jira/config:config",
    ],
    visibility = [
        "PUBLIC"
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
