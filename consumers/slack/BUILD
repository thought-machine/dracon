subinclude("//third_party/defs:docker")


go_binary(
    name = "slack",
    srcs = [
        "main.go",
    ],
    deps = [
        "//consumers",
        "//consumers/slack/utils:utils",
        "//api/proto:v1",
        "//third_party/go:protobuf",
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
