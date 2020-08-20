go_library(
    name = "utils",
    srcs = [
        "utils.go",
    ],
    deps = [
        "//consumers",
        "//consumers/jira_c/document:document",
        "//api/proto:v1",
        "//third_party/go:protobuf",
        ],
    visibility = [
        "//consumers/jira_c:jira_c",
    ],
)

go_test(
    name = "utils_test",
    srcs = [
        "utils_test.go",

    ],
    deps = [
        ":utils",
        "//consumers",
        "//api/proto:v1",
        "//third_party/go:gogo_protobuf",
        "//third_party/go:protobuf",
        "//third_party/go:stretchr_testify",
    ],
)
