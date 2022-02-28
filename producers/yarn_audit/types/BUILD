go_library(
    name = "types",
    srcs = [
        "yarn-issue.go",
    ],
    visibility = ["//producers/yarn_audit/..."],
    deps = [
        "//api/proto:v1",
        "//producers",
    ],
)

go_test(
    name = "types_test",
    srcs = [
        "yarn-issue_test.go",
    ],
    deps = [
        ":types",
        "//api/proto:v1",
        "//producers",
        "//third_party/go:stretchr_testify",
    ],
)
