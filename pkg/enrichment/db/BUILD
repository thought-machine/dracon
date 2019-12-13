go_library(
    name = "db",
    srcs = [
        "db.go",
        "issue.go",
    ],
    visibility = [
        "//cmd/enricher",
        "//pkg/enrichment",
    ],
    deps = [
        "//pkg/enrichment/db/migrations",
        "//pkg/genproto/v1",
        "//third_party/go:golang-migrate_migrate",
        "//third_party/go:golang_protobuf",
        "//third_party/go:jmoiron_sqlx",
        "//third_party/go:rakyll_statik",
    ],
)
