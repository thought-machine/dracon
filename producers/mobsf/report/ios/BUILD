go_library(
    name = "ios",
    srcs = [
        "ios.go",
    ],
    deps = [
        "//api/proto:v1",
        "//producers/mobsf/report:report",
    ],
    visibility = ["//producers/mobsf/..."]
)

go_test(
    name = "ios_test",
    srcs = [
        "ios_test.go",
    ],
    deps = [
        ":ios",
        "//api/proto:v1",
        "//producers:producers",
        "//third_party/go:stretchr_testify",
    ],
)
