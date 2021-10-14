go_library(
    name = "npmfullaudit",
    srcs = [
        "npmfullaudit.go",
    ],
    visibility = ["//producers/npm_audit/..."],
    deps = [
        "//api/proto:v1",
        "//producers",
        "//producers/npm_audit/types",
    ],
)

go_test(
    name = "npmfullaudit_test",
    srcs = [
        "npmfullaudit_test.go",
    ],
    deps = [
        ":npmfullaudit",
        "//api/proto:v1",
        "//producers",
        "//third_party/go:stretchr_testify",
    ],
)
