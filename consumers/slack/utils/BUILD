go_library(
    name = "utils",
    srcs = [
        "utils.go",
    ],
    visibility = [
        "//consumers/slack/...",
    ],
    deps = [
        "//api/proto:v1",
        "//consumers/slack/types",
        "//third_party/go:protobuf",
    ],
)

go_test(
    name = "slack_go_test",
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
