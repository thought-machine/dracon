go_library(
    name = "consumers",
    srcs = [
        "consumer.go",
    ],
    visibility = ["//consumers/..."],
    deps = [
        "//api/proto:v1",
        "//pkg/putil",
        "//third_party/go:gogo_protobuf",
        "//third_party/go:protobuf",
    ],
)

go_test(
    name = "consumers_go_test",
    srcs = [
        "consumer.go",
        "consumer_test.go",
    ],
    deps = [
        "//api/proto:v1",
        "//pkg/putil",
        "//third_party/go:gogo_protobuf",
        "//third_party/go:protobuf",
        "//third_party/go:stretchr_testify",
    ],
)

python_test(
    name = "consumers_py_test",
    srcs = [
        "consumer_test.py",
    ],
    deps = [
        ":consumers_base_python",
        "//api/proto:v1",
        "//third_party/python:protobuf",
    ],
)

python_binary(
    name = "consumers_base_python",
    main = "consumer.py",
    visibility = ["//consumers/..."],
    deps = [
        "//third_party/python:protobuf",
    ],
)
