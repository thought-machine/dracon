subinclude("//third_party/defs:docker")

go_binary(
    name = "elasticsearch",
    srcs = [
        "main.go",
    ],
    static = True,
    deps = [
        "//api/proto:v1",
        "//consumers",
        "//third_party/go:elastic_go-elasticsearch_v8",
        "//third_party/go:protobuf",
    ],
)

go_test(
    name = "elasticsearch_test",
    srcs = [
        "main.go",
        "main_test.go",
    ],
    deps = [
        "//api/proto:v1",
        "//consumers",
        "//third_party/go:elastic_go-elasticsearch_v8",
        "//third_party/go:gogo_protobuf",
        "//third_party/go:protobuf",
        "//third_party/go:stretchr_testify",
    ],
)

docker_image(
    name = "dracon-consumer-elasticsearch",
    srcs = [
        ":elasticsearch",
    ],
    base_image = "//build/docker:dracon-base-go",
    image = "dracon-consumer-elasticsearch",
    visibility = ["//examples/..."],
)
