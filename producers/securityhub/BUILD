subinclude("//third_party/defs:docker")

go_binary(
    name = "securityhub",
    srcs = [
        "main.go",
    ],
    deps = [
        "//api/proto:v1",
        "//producers",
        "//third_party/go:aws-sdk-go-v2-securityhub",
    ]
)

go_test(
    name = "securityhub_test",
    srcs = [
        "main.go",
        "main_test.go",
    ],
    data = glob(["./testcases/*.json"]),
    deps = [
        "//api/proto:v1",
        "//producers",
        "//pkg/putil",
        "//third_party/go:aws-sdk-go-v2-securityhub",
        "//third_party/go:stretchr_testify",
    ]
)

docker_image(
    name = "dracon-producer-securityhub",
    srcs = [
        ":securityhub",
    ],
    base_image = "//build/docker:dracon-base-go",
    image = "dracon-producer-securityhub",
    visibility = ["//examples/..."],
)
