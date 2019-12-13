subinclude("@pleasings//docker")

go_binary(
    name = "java_spotbugs",
    srcs = [
        "main.go",
    ],
    deps = [
        "//pkg/genproto/v1",
        "//producers",
    ],
)

docker_image(
    name = "dracon-producer-java_spotbugs",
    srcs = [
        ":java_spotbugs",
    ],
    base_image = "//build/docker:dracon-base-go",
    image = "dracon-producer-java_spotbugs",
)
