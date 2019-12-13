subinclude("@pleasings//docker")

go_binary(
    name = "enricher",
    srcs = [
        "main.go",
    ],
    static = True,
    deps = [
        "//pkg/enrichment",
        "//pkg/enrichment/db",
        "//pkg/genproto/v1",
        "//pkg/putil",
        "//third_party/go:spf13_cobra",
        "//third_party/go:spf13_viper",
    ],
)

docker_image(
    name = "dracon-enricher",
    srcs = [":enricher"],
    base_image = "//build/docker:dracon-base-go",
    image = "dracon-enricher",
)
