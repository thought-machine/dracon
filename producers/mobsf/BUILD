subinclude("//third_party/defs:docker")

go_binary(
    name = "entrypoint",
    srcs = [
        "cli.go",
        "main.go",
        "project.go",
    ],
    out = "entrypoint",
    deps = [
        "//api/proto:v1",
        "//pkg/template",
        "//producers",
        "//producers/mobsf/report",
        "//producers/mobsf/report/android",
        "//producers/mobsf/report/ios",
    ],
)

docker_image(
    name = "mobsf",
    srcs = [
        ":entrypoint",
    ],
    dockerfile = "Dockerfile-producer-mobsf",
    image = "dracon-producer-mobsf",
)
