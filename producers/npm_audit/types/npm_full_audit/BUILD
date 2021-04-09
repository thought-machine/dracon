go_library(
    name = "npm_full_audit",
    srcs = [
        "npm_full_audit.go",
    ],
    visibility = ["//producers/npm_audit/..."],
    deps = [
        "//api/proto:v1",
        "//producers",
        "//producers/npm_audit/types",
    ],
)

go_test(
    name = "npm_full_audit_test",
    srcs = [
        "npm_full_audit_test.go",
    ],
    deps = [
        ":npm_full_audit",
        "//api/proto:v1",
        "//producers",
        "//third_party/go:stretchr_testify",
    ],
)
