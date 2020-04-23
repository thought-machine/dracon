go_library(
    name = "producers",
    srcs = [
        "producer.go",
    ],
    visibility = ["//producers/..."],
    deps = [
        "//api/proto:v1",
        "//pkg/putil",
        "//third_party/go:protobuf",
    ],
)

go_test(
    name = "producers_go_test",
    srcs = [
        "producer.go",
        "producer_test.go",
    ],
    deps = [
        "//api/proto:v1",
        "//pkg/putil",
        "//third_party/go:gogo_protobuf",
        "//third_party/go:protobuf",
        "//third_party/go:stretchr_testify",
    ],
)

# TODO(hjenkins) support building python
