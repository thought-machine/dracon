go_library(
    name = "db",
    srcs = [
        "db.go",
        "issue.go",
    ],
    resources = [":resources"],
    visibility = [
        "//cmd/enricher",
        "//pkg/enrichment/...",
        "//producers/...",
    ],
    deps = [
        "//api/proto:v1",
        "//third_party/go:golang-migrate_migrate",
        "//third_party/go:jmoiron_sqlx",
        "//third_party/go:lib_pq",
        "//third_party/go:protobuf",
    ],
)

# This is a workaround for embedding files that are
# not in the same pkg dir.
filegroup(
    name = "resources",
    srcs = ["//enrichment_service/configs/sql/migrations"],
)
