subinclude("@pleasings//docker")


go_binary(
    name = "slack",
    srcs = [
        "main.go",
    ],
    deps = [
        "//consumers",
        "//consumers/slack/utils:utils",
        "//pkg/genproto/v1",
        "//third_party/go:golang_protobuf",
    ],
)


docker_image(
    name = "image",
    srcs = [
        ":slack",
    ],
    base_image = "//build/docker:dracon-base-go",
    image = "dracon-consumer-slack",
)
