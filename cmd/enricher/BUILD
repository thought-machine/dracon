subinclude("//third_party/defs:docker")

go_binary(
    name = "enricher",
    srcs = [
        "main.go",
    ],
    static = True,
    deps = [
        "//api/proto:v1",
        "//pkg/enrichment",
        "//pkg/enrichment/db",
        "//pkg/putil",
        "//third_party/go:protobuf",
        "//third_party/go:spf13_cobra",
        "//third_party/go:spf13_viper",
    ],
)

docker_image(
    name = "dracon-enricher",
    srcs = [":enricher"],
    base_image = "//build/docker:dracon-base-go",
    image = "dracon-enricher",
    visibility = ["//examples/..."],
)
