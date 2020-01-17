go_library(
    name = "kubernetes",
    srcs = [
        "apply.go",
    ],
    visibility = ["PUBLIC"],
    deps = [
    ],
)

go_test(
    name = "kubernetes_test",
    srcs = [
        "apply_test.go",
    ],
    deps = [
        ":kubernetes",
        "//third_party/go:stretchr_testify",
    ],
)
