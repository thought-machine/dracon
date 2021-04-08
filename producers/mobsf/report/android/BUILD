go_library(
    name = "android",
    srcs = [
        "android.go",
    ],
    visibility = ["//producers/mobsf/..."],
    deps = [
        "//api/proto:v1",
        "//producers/mobsf/report",
    ],
)

go_test(
    name = "android_test",
    srcs = [
        "android_test.go",
    ],
    deps = [
        ":android",
        "//api/proto:v1",
        "//producers",
        "//third_party/go:stretchr_testify",
    ],
)
