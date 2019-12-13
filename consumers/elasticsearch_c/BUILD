subinclude("@pleasings//docker")

go_binary(
    name = "elasticsearch_c",
    srcs = [
        "main.go",
    ],
    static = True,
    deps = [
        "//consumers",
        "//pkg/genproto/v1",
        "//third_party/go:elastic_go-elasticsearch_v6",
        "//third_party/go:golang_protobuf",
    ],
)

docker_image(
    name = "dracon-consumer-elasticsearch",
    srcs = [
        ":elasticsearch_c",
    ],
    base_image = "//build/docker:dracon-base-go",
    image = "dracon-consumer-elasticsearch",
)
