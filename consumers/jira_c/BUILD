subinclude("@third_party/subrepos/pleasings//docker")

go_binary(
    name = "jira_c",
    srcs = [
        "main.go",
    ],
    deps = [
        "//consumers",
        "//consumers/jira_c/utils:utils",
        "//consumers/jira_c/gojira:gojira",
        "//api/proto:v1",
        "//third_party/go:protobuf",
    ],
)

filegroup(
    name = "config",
    srcs = ["config.yaml"],
    visibility = ["//consumers/jira_c/..."],
)

docker_image(
    name = "image",
    srcs = [
        ":jira_c",
        ":config",
    ],
    base_image = "//build/docker:dracon-base-go",
    image = "dracon-consumer-jira",
)
