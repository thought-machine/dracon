subinclude("//third_party/defs:docker")

go_binary(
    name = "golang_gosec",
    srcs = [
        "main.go",
    ],
    deps = [
        "//api/proto:v1",
        "//producers",
    ],
)

go_test(
    name = "golang_gosec_test",
    srcs = [
        "main.go",
        "main_test.go",
    ],
    deps = [
        "//api/proto:v1",
        "//producers",
        "//third_party/go:stretchr_testify",
    ],
)

docker_image(
    name = "dracon-producer-gosec",
    srcs = [
        ":golang_gosec",
    ],
    base_image = "//build/docker:dracon-base-go",
    image = "dracon-producer-gosec",
    visibility = ["//examples/..."],
)
