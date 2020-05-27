subinclude("@third_party/subrepos/pleasings//docker")

go_binary(
    name = "python_bandit",
    srcs = [
        "main.go",
    ],
    deps = [
        "//api/proto:v1",
        "//producers",
    ],
)

go_test(
    name = "python_bandit_test",
    srcs = [
        "main_test.go",
    ],
)

docker_image(
    name = "dracon-producer-bandit",
    srcs = [
        ":python_bandit",
    ],
    base_image = "//build/docker:dracon-base-go",
    image = "dracon-producer-bandit",
)
