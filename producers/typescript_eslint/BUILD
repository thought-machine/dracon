subinclude("//third_party/defs:docker")

go_binary(
    name = "typescript_eslint",
    srcs = [
        "main.go",
    ],
    deps = [
        "//api/proto:v1",
        "//producers",
        "//producers/typescript_eslint/types",
    ],
)

go_test(
    name = "typescript_eslint_test",
    srcs = [
        "main.go",
        "main_test.go",
    ],
    deps = [
        "//api/proto:v1",
        "//producers",
        "//producers/typescript_eslint/types",
        "//third_party/go:stretchr_testify",
    ],
)

docker_image(
    name = "dracon-producer-eslint",
    srcs = [
        ":typescript_eslint",
    ],
    base_image = "//build/docker:dracon-base-go",
    image = "dracon-producer-eslint",
)
