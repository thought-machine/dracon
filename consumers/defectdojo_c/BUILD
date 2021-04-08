subinclude("//third_party/defs:docker")

python_binary(
    name = "defectdojo",
    main = "defectdojo_c.py",
    visibility = ["//security/dracon/..."],
    deps = [
        "//api/proto:v1",
        "//consumers:consumers_base_python",
        "//third_party/python:defectdojo_api",
        "//third_party/python:protobuf",
    ],
)

python_test(
    name = "defectdojo_test",
    srcs = ["defectdojo_test.py"],
    deps = [
        ":defectdojo",
        "//api/proto:v1",
        "//consumers:consumers_base_python",
        "//third_party/python:protobuf",
    ],
)

docker_image(
    name = "dracon_consumer_defectdojo",
    srcs = [
        ":defectdojo",
    ],
    base_image = "//build/docker:dracon-base-python",
    image = "dracon-consumer-defectdojo",
)
