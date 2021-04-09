subinclude("//third_party/defs:docker")

# this producer covers nancy https://github.com/sonatype-nexus-community/nancy

go_binary(
    name = "nancy",
    srcs = [
        "main.go",
    ],
    deps = [
        "//api/proto:v1",
        "//producers",
        "//producers/golang_nancy/types:nancy-issue",
    ],
)

go_test(
    name = "nancy_test",
    srcs = [
        "main.go",
        "main_test.go",
    ],
    deps = [
        "//api/proto:v1",
        "//producers",
        "//producers/golang_nancy/types:nancy-issue",
        "//third_party/go:stretchr_testify",
    ],
)

docker_image(
    name = "dracon-producer-nancy",
    srcs = [
        ":nancy",
    ],
    base_image = "//build/docker:dracon-base-go",
    image = "dracon-producer-nancy",
)
