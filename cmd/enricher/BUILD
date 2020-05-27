subinclude("@third_party/subrepos/pleasings//docker")

go_binary(
    name = "enricher",
    srcs = [
        "main.go",
    ],
    static = True,
    deps = [
        "//pkg/enrichment",
        "//pkg/enrichment/db",
        "//api/proto:v1",
        "//pkg/putil",
        "//third_party/go:spf13_cobra",
        "//third_party/go:spf13_viper",
        "//third_party/go:protobuf",
    ],
)

docker_image(
    name = "dracon-enricher",
    srcs = [":enricher"],
    base_image = "//build/docker:dracon-base-go",
    image = "dracon-enricher",
)
