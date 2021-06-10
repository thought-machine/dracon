subinclude("//third_party/defs:docker")

go_binary(
    name = "zap",
    srcs = [
        "main.go",
    ],
    deps = [
        "//api/proto:v1",
        "//producers",
        "//producers/zap_producer/types",
    ],
)

go_test(
    name = "zap_test",
    srcs = [
        "main.go",
        "main_test.go",
    ],
    deps = [
        "//api/proto:v1",
        "//producers",
        "//producers/zap_producer/types",
        "//third_party/go:stretchr_testify",
    ],
)

docker_image(
    name = "dracon-producer-zap",
    srcs = [
        ":zap",
    ],
    base_image = "//build/docker:dracon-base-go",
    image = "dracon-producer-zap",
)
