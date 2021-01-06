subinclude("@third_party/subrepos/pleasings//docker")

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
        "//pkg/template:template",
        "//producers:producers",
        "//producers/mobsf/report:report",
        "//producers/mobsf/report/android:android",
        "//producers/mobsf/report/ios:ios",
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
