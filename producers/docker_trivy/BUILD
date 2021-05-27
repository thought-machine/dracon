subinclude("//third_party/defs:docker")

# this producer covers nancy https://github.com/sonatype-nexus-community/nancy

go_binary(
    name = "trivy",
    srcs = [
        "main.go",
    ],
    deps = [
        "//api/proto:v1",
        "//producers",
        "//producers/docker_trivy/types",
    ],
)

go_test(
    name = "trivy_test",
    srcs = [
        "main.go",
        "main_test.go",
    ],
    deps = [
        "//api/proto:v1",
        "//producers",
        "//producers/docker_trivy/types",
        "//third_party/go:stretchr_testify",
    ],
)

docker_image(
    name = "dracon-producer-trivy",
    srcs = [
        ":trivy",
    ],
    base_image = "//build/docker:dracon-base-go",
    image = "dracon-producer-trivy",
)
