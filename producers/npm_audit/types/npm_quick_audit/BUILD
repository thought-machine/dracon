go_library(
    name = "npm_quick_audit",
    srcs = [
        "npm_advisory.go",
        "npm_quick_audit.go",
    ],
    deps = [
        "//api/proto:v1",
        "//producers",
        "//producers/npm_audit/types:types",
    ],
    visibility = ["//producers/npm_audit/..."]
)

go_test(
    name = "npm_quick_audit_test",
    srcs = [
        "npm_advisory_test.go",
        "npm_quick_audit_test.go",
    ],
    data = [
        ":npm_quick_audit_test_data",
    ],
    deps = [
        ":npm_quick_audit",
        "//api/proto:v1",
        "//producers",
        "//third_party/go:h2non_gock",
        "//third_party/go:stretchr_testify",
    ],
)

filegroup(
    name = "npm_quick_audit_test_data",
    srcs = [
        "npm_advisory_1556",
        "npm_advisory_no_advisorydata",
        "npm_advisory_not_json",
    ],
    visibility = [":npm_quick_audit_test"],
)
