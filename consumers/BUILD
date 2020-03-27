go_library(
    name = "consumers",
    srcs = [
        "consumer.go",
    ],
    visibility = ["//consumers/..."],
    deps = [
        "//pkg/genproto/v1",
        "//pkg/putil",
        "//third_party/go:gogo_protobuf",
        "//third_party/go:golang_protobuf",
    ],
)

go_test(
    name = "consumers_go_test",
    srcs = [
        "consumer.go",
        "consumer_test.go",
    ],
    deps = [
        "//pkg/genproto/v1",
        "//pkg/putil",
        "//third_party/go:gogo_protobuf",
        "//third_party/go:golang_protobuf",
        "//third_party/go:stretchr_testify",
    ],
)
python_test(
    name = "consumers_py_test",
    srcs = [
            "consumer_test.py"
            ],
    deps = [
            "//third_party/python:protobuf",
            "//api/proto:engine",
            ":consumers_base_python",
            ]
)

python_binary(
    name = "consumers_base_python",
    main = "consumer.py",
    visibility = ["//consumers/..."],
    deps = [
            "//third_party/python:protobuf"
            ],
)
