subinclude("//third_party/defs:go_mock")

go_mock(
    name = "mock",
    interfaces = [
        "EnrichDatabase",
    ],
    package = "github.com/thought-machine/dracon/pkg/enrichment/db",
    src_lib = "//pkg/enrichment/db",
    visibility = [
        "//producers/...",
    ],
    deps = [
        "//api/proto:v1",
    ],
)
