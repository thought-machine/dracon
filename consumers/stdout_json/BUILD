subinclude("@pleasings//docker")

python_binary(
    name = "stdout_json",
    main = "stdout_json.py",
    deps = [
        "//consumers:consumers_base_python",
        "//api/proto:config",
        "//api/proto:engine",
        "//api/proto:issue",
    ],
)
python_test(
    name = "stdout_json_test",
    srcs = [
            "stdout_json_test.py"
            ],
    deps = [
            "//third_party/python:protobuf",
            "//api/proto:engine",
            "//consumers:consumers_base_python",
            ":stdout_json",
            ]
)
docker_image(
    name = "dracon-consumer-stdout",
    srcs = [
        ":stdout_json",
    ],
    base_image = "//build/docker:dracon-base-python",
    image = "dracon-consumer-stdout-json",
)
