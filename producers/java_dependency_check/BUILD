go_binary(
    name = "dependency_check",
    srcs = [
        "main.go",
    ],
    deps = [
        "//api/proto:v1",
        "//producers",
    ],
)

go_test(
    name = "dependency_check_test",
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
