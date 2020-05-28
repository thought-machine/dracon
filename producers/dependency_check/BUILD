subinclude("@pleasings//docker")

go_binary(
    name = "dependency_check",
    srcs = [
        "main.go",
    ],
    deps = [
        "//api/proto:v1",
        "//producers",
    ],
)

go_test(
    name = "dependency_check_test",
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
    name = "dracon-producer-dependency-check",
    srcs = [
        ":dependency_check",
    ],
    base_image = "//build/docker:dracon-base-go",
    image = "dracon-producer-dependency-check",
)
