subinclude("@third_party/subrepos/pleasings//go:mock")

go_library(
    name = "db",
    srcs = [
        "db.go",
        "issue.go",
    ],
    visibility = [
        "//cmd/enricher",
        "//pkg/enrichment/...",
        "//producers/..."
    ],
    deps = [
        "//pkg/enrichment/db/migrations",
        "//api/proto:v1",
        "//third_party/go:golang-migrate_migrate",
        "//third_party/go:protobuf",
        "//third_party/go:jmoiron_sqlx",
        "//third_party/go:lib_pq",
        "//third_party/go:rakyll_statik",
    ],
)


go_mock(
    name = "db_mock",
    interfaces = [
        "EnrichDatabase",
    ],
    # package = "pkg/enrichment/db/db",
    src_lib = ":db",
    visibility = [
        "PUBLIC", # todo change
    ],
    deps = [
        "//api/proto:v1",
    ],
)
