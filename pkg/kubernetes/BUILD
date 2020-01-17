go_library(
    name = "kubernetes",
    srcs = [
        "apply.go",
    ],
    visibility = ["PUBLIC"],
    deps = [
        "//third_party/go:pkg_errors",
    ],
)

go_test(
    name = "kubernetes_test",
    srcs = [
        "apply_test.go",
    ],
    deps = [
        ":kubernetes",
        "//third_party/go:pkg_errors",
        "//third_party/go:stretchr_testify",
    ],
)
