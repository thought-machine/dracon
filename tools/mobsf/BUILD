subinclude("@third_party/subrepos/pleasings//docker")

go_binary(
    name = "entrypoint",
    srcs = [
        "main.go",
    ],
    out = "entrypoint",
    deps = [
        "//pkg/template:template",
    ],
)

docker_image(
    name = "mobsf",
    srcs = [
        ":entrypoint",
    ],
    dockerfile = "Dockerfile-tool-mobsf",
    image = "dracon-tool-mobsf",
)
