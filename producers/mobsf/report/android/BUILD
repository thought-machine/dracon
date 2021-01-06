go_library(
    name = "android",
    srcs = [
        "android.go",
    ],
    deps = [
        "//api/proto:v1",
        "//producers/mobsf/report:report",
    ],
    visibility = ["//producers/mobsf/..."]
)

go_test(
    name = "android_test",
    srcs = [
        "android_test.go",
    ],
    deps = [
        ":android",
        "//api/proto:v1",
        "//producers:producers",
        "//third_party/go:stretchr_testify",
    ],
)
