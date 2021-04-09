go_library(
    name = "enrichment",
    srcs = [
        "issue.go",
    ],
    visibility = ["PUBLIC"],
    deps = [
        "//api/proto:v1",
        "//pkg/enrichment/db",
        "//third_party/go:protobuf",
    ],
)

go_test(
    name = "enrichment_test",
    srcs = [
        "issue.go",
        "issue_test.go",
    ],
    deps = [
        "//api/proto:v1",
        "//pkg/enrichment/db",
        "//third_party/go:protobuf",
        "//third_party/go:stretchr_testify",
    ],
)
