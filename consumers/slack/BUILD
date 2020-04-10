subinclude("@pleasings//docker")

go_binary(
    name = "slack_c",
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

docker_image(
    name = "dracon-consumer-slack",
    srcs = [
        ":slack_c",
    ],
    base_image = "//build/docker:dracon-base-go",
    image = "dracon-consumer-slack",
)
