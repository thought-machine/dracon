go_library(
    name = "utils",
    srcs = [
        "utils.go",
    ],
    visibility = [
        "//consumers/jira_c/...",
    ],
    deps = [
        "//consumers",
        "//api/proto:v1",
        "//third_party/go:protobuf",
        ],
)

go_test(
    name = "utils_test",
    srcs = [
        "utils_test.go",

    ],
    visibility = ["//consumers/jira_c/..."],
    deps = [
        ":utils",
        "//consumers",
        "//api/proto:v1",
        "//third_party/go:gogo_protobuf",
        "//third_party/go:protobuf",
        "//third_party/go:stretchr_testify",
    ],
)
