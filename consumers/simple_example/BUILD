subinclude("@third_party/subrepos/pleasings//docker")

python_binary(
    name = "simple",
    main = "simple.py",
    deps = [
        "//third_party/python:requests",
        "//consumers:consumers_base_python",
        "//third_party/python:protobuf",
        "//api/proto:v1",
    ],
)

python_test(
    name = "simple_test",
    srcs = ["simple_test.py"],
    deps = [
        ":simple",
        "//api/proto:v1",
        "//consumers:consumers_base_python",
        "//third_party/python:protobuf",
        "//third_party/python:requests",
    ],
)

docker_image(
    name = "dracon-consumer-simple",
    srcs = [
        ":simple",
    ],
    base_image = "//build/docker:dracon-base-python",
    image = "dracon-consumer-simple",
)