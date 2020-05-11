subinclude("@pleasings//docker")

go_binary(
    name = "slack",
    srcs = [
        "main.go",
    ],
    static = True,
    deps = [
        "//consumers",
        "//pkg/genproto/v1",
        "//third_party/go:golang_protobuf",
    ],
)

go_test(
    name = "slack_go_test",
    srcs = [
        "main.go",
        "main_test.go",
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
    name = "image",
    srcs = [
        ":slack_c",
    ],
    base_image = "//build/docker:dracon-base-go",
    image = "dracon-consumer-slack",
)
