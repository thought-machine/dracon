subinclude("//third_party/defs:docker")

# this producer covers nancy https://github.com/sonatype-nexus-community/nancy

go_binary(
    name = "yarn_audit",
    srcs = [
        "main.go",
    ],
    deps = [
        "//api/proto:v1",
        "//producers",
        "//producers/yarn_audit/types",
    ],
)

docker_image(
    name = "dracon-producer-yarn-audit",
    srcs = [
        ":yarn_audit",
    ],
    base_image = "//build/docker:dracon-base-go",
    image = "dracon-producer-yarn-audit",
)
