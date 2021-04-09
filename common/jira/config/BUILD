go_library(
    name = "config",
    srcs = [
        "config.go",
        "types.go",
    ],
    visibility = [
        "PUBLIC",
    ],
    deps = [
        "//third_party/go:yaml.v2",
    ],
)

go_test(
    name = "config_test",
    srcs = [
        "config_test.go",
    ],
    deps = [
        ":config",
        "//third_party/go:stretchr_testify",
        "//third_party/go:yaml.v2",
    ],
)
