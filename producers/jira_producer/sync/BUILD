
go_library(
    name = "sync",
    srcs = [
        "sync.go",
    ],
    visibility = ["//producers/..."],
    deps = [
        "//api/proto:v1",
        "//common/jira/document",
        "//pkg/enrichment",
        "//third_party/go:go-jira",
        "//third_party/go:mitchellh_mapstructure",
        "//common/jira/config",
        "//pkg/enrichment/db",
        "//third_party/go:tgo",
    ],
)

go_test(
    name = "sync_test",
    srcs = [
        "sync_test.go", "sync.go",
    ],
    deps = [
        "//third_party/go:go-jira",
        "//third_party/go:stretchr_testify",
          "//api/proto:v1",
        "//common/jira/document",
        "//pkg/enrichment",
        # "//pkg/enrichment/db",
        "//third_party/go:go-jira",
        "//third_party/go:mitchellh_mapstructure",
        "//common/jira/config",
        "//third_party/go:tgo",
        "//pkg/enrichment/db:db_mock"
    ],
)
