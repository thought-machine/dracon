
go_library(
    name = "types",
    srcs = [
        "safety-issue.go",
    ],
    visibility = ["//producers/pipsafety/..."]

)
go_test(
    name = "types-test",
    srcs = [
            "safety-issue.go",
            "safety-issue_test.go"
    ],
    deps = [
        "//third_party/go:stretchr_testify",
    ],
)
