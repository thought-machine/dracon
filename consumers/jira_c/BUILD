subinclude("//third_party/defs:docker")

go_binary(
    name = "jira_c",
    srcs = [
        "main.go",
    ],
    deps = [
        "//consumers",
        "//common/jira/config:config",
        "//consumers/jira_c/utils:utils",
        "//common/jira/jira:jira",
        "//api/proto:v1",
        "//third_party/go:protobuf",
        "//third_party/go:yaml.v2",
    ],
)


docker_image(
    name = "image",
    srcs = [
        ":jira_c",
        "//common/jira:config_file",
    ],
    base_image = "//build/docker:dracon-base-go",
    image = "dracon-consumer-jira",
)
