subinclude("@third_party/subrepos/pleasings//docker")

go_binary(
    name = "elasticsearch_c",
    srcs = [
        "main.go",
    ],
    static = True,
    deps = [
        "//consumers",
        "//api/proto:v1",
        "//third_party/go:elastic_go-elasticsearch_v7",
        "//third_party/go:protobuf",
    ],
)

go_test(
    name = "elasticsearch_c_test",
    srcs = [
        "main_test.go",
        "main.go"

    ],
    deps = [
        "//consumers",
        "//pkg/genproto/v1",
        "//third_party/go:gogo_protobuf",
        "//third_party/go:golang_protobuf",
        "//third_party/go:stretchr_testify",
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
