go_library(
    name = "utils",
    srcs = [
        "utils.go",
    ],
    visibility = [
        "//consumers/jira_c",
    ],
    deps = [
        "//api/proto:v1",
        "//common/jira/document",
        "//consumers",
        "//third_party/go:protobuf",
    ],
)

go_test(
    name = "utils_test",
    srcs = [
        "utils_test.go",
    ],
    deps = [
        ":utils",
        "//api/proto:v1",
        "//consumers",
        "//third_party/go:gogo_protobuf",
        "//third_party/go:protobuf",
        "//third_party/go:stretchr_testify",
    ],
)
