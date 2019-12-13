go_library(
    name = "producers",
    srcs = [
        "producer.go",
    ],
    visibility = ["//producers/..."],
    deps = [
        "//pkg/genproto/v1",
        "//pkg/putil",
        "//third_party/go:golang_protobuf",
    ],
)

go_test(
    name = "producers_go_test",
    srcs = [
        "producer.go",
        "producer_test.go",
    ],
    deps = [
        "//pkg/genproto/v1",
        "//pkg/putil",
        "//third_party/go:gogo_protobuf",
        "//third_party/go:golang_protobuf",
        "//third_party/go:stretchr_testify",
    ],
)

# TODO(hjenkins) support building python
