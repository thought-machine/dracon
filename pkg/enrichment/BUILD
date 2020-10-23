go_library(
    name = "enrichment",
    srcs = [
        "issue.go",
    ],
    visibility = ["//cmd/enricher"],
    deps = [
        "//pkg/enrichment/db",
        "//api/proto:v1",
        "//third_party/go:protobuf",
    ],
)

go_test(
    name = "enrichment_test",
    srcs = [
        "issue.go",
        "issue_test.go"
    ],
    deps = [
        "//api/proto:v1",
        "//third_party/go:stretchr_testify",
        "//pkg/enrichment/db",
        "//api/proto:v1",
        "//third_party/go:protobuf",
    ],
)
