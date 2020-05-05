subinclude("@pleasings//docker")

go_binary(
    name = "python_pipsafety",
    srcs = [
        "main.go",
    ],
    deps = [
        "//api/proto:v1",
        "//producers",
    ],
)

go_test(
    name = "python_pipsafety-test",
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
    name = "dracon-producer-pipsafety",
    srcs = [
        ":python_pipsafety",
    ],
    base_image = "//build/docker:dracon-base-go",
    image = "dracon-producer-pipsafety",
)
