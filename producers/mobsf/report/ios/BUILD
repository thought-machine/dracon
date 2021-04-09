go_library(
    name = "ios",
    srcs = [
        "ios.go",
    ],
    visibility = ["//producers/mobsf/..."],
    deps = [
        "//api/proto:v1",
        "//producers/mobsf/report",
    ],
)

go_test(
    name = "ios_test",
    srcs = [
        "ios_test.go",
    ],
    deps = [
        ":ios",
        "//api/proto:v1",
        "//producers",
        "//third_party/go:stretchr_testify",
    ],
)
