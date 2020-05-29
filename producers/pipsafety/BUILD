subinclude("@third_party/subrepos/pleasings//docker")

go_binary(
    name = "pipsafety",
    srcs = [
        "main.go",
    ],
    deps = [
        "//api/proto:v1",
        "//producers",
        "//producers/pipsafety/types",
    ],
)

go_test(
    name = "pipsafety-test",
    srcs = [
            "main.go",
            "main_test.go"
    ],
    deps = [
        "//api/proto:v1",
        "//producers",
        "//producers/pipsafety/types",
        "//third_party/go:stretchr_testify",
    ],
)

docker_image(
    name = "image",
    srcs = [
        ":pipsafety",
    ],
    base_image = "//build/docker:dracon-base-go",
    image = "dracon-producer-pipsafety",
)
