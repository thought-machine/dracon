go_library(
    name = "npmquickaudit",
    srcs = [
        "npm_advisory.go",
        "npmquickaudit.go",
    ],
    visibility = ["//producers/npm_audit/..."],
    deps = [
        "//api/proto:v1",
        "//producers",
        "//producers/npm_audit/types",
    ],
)

go_test(
    name = "npmquickaudit_test",
    srcs = [
        "npm_advisory_test.go",
        "npmquickaudit_test.go",
    ],
    data = [
        ":npmquickaudit_test_data",
    ],
    deps = [
        ":npmquickaudit",
        "//api/proto:v1",
        "//producers",
        "//third_party/go:h2non_gock",
        "//third_party/go:stretchr_testify",
    ],
)

filegroup(
    name = "npmquickaudit_test_data",
    srcs = [
        "npm_advisory_1556",
        "npm_advisory_no_advisorydata",
        "npm_advisory_not_json",
    ],
    visibility = [":npmquickaudit_test"],
)
