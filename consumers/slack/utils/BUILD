
go_library(
    name = "utils",
    srcs = [
        "utils.go",
    ],
    visibility = [
        "//consumers/slack/...",
    ],
    deps = ["//pkg/genproto/v1",
         "//consumers/slack/types",
        "//third_party/go:golang_protobuf",]
        
)
go_test(
    name = "slack_go_test",
    srcs = [
        "utils_test.go",

    ],
    deps = [
        ":utils",
        "//consumers",
        "//pkg/genproto/v1",
        "//third_party/go:gogo_protobuf",
        "//third_party/go:golang_protobuf",
        "//third_party/go:stretchr_testify",
    ],
)
