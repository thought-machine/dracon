go_library(
    name = "sync",
    srcs = [
        "sync.go",
    ],
    visibility = ["//producers/..."],
    deps = [
        "//api/proto:v1",
        "//common/jira/config",
        "//common/jira/document",
        "//pkg/enrichment",
        "//pkg/enrichment/db",
        "//third_party/go:go-jira",
        "//third_party/go:mitchellh_mapstructure",
        "//third_party/go:tgo",
    ],
)

go_test(
    name = "sync_test",
    srcs = [
        "sync.go",
        "sync_test.go",
    ],
    deps = [
        "//api/proto:v1",
        "//common/jira/config",
        "//common/jira/document",
        "//pkg/enrichment",
        "//pkg/enrichment/db",
        "//pkg/enrichment/db/mock",
        "//third_party/go:go-jira",
        "//third_party/go:mitchellh_mapstructure",
        "//third_party/go:mock",
        "//third_party/go:stretchr_testify",
        "//third_party/go:tgo",
    ],
)
