subinclude("@third_party/subrepos/pleasings//docker")

go_binary(
    name = "jira_c",
    srcs = [
        "main.go",
    ],
    deps = [
        "//consumers",
        "//consumers/jira_c/config:config",
        "//consumers/jira_c/utils:utils",
        "//consumers/jira_c/jira:jira",
        "//api/proto:v1",
        "//third_party/go:protobuf",
        "//third_party/go:yaml.v2",
    ],
)

filegroup(
    name = "config_file",
    srcs = ["config.yaml"],
)

docker_image(
    name = "image",
    srcs = [
        ":jira_c",
        ":config_file",
    ],
    base_image = "//build/docker:dracon-base-go",
    image = "dracon-consumer-jira",
)
