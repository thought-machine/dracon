subinclude("//third_party/defs:docker")

go_binary(
    name = "java_spotbugs",
    srcs = [
        "main.go",
    ],
    deps = [
        "//api/proto:v1",
        "//producers",
    ],
)

go_test(
    name = "java_spotbugs_test",
    srcs = [
        "main.go",
        "main_test.go",
    ],
    deps = [
        "//api/proto:v1",
        "//producers",
        "//third_party/go:protobuf",
        "//third_party/go:stretchr_testify",
    ],
)

docker_image(
    name = "dracon-producer-java_spotbugs",
    srcs = [
        ":java_spotbugs",
    ],
    base_image = "//build/docker:dracon-base-go",
    image = "dracon-producer-java_spotbugs",
    visibility = ["//examples/..."],
)
