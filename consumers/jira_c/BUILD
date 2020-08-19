subinclude("@third_party/subrepos/pleasings//docker")

go_binary(
    name = "jira_c",
    srcs = [
        "main.go",
    ],
    deps = [
        ":configuration",
        "//consumers",
        "//consumers/jira_c/types:config",
        "//consumers/jira_c/utils:utils",
        "//consumers/jira_c/jira:jira",
        "//api/proto:v1",
        "//third_party/go:protobuf",
        "//third_party/go:yaml.v2",
    ],
)

go_library(
    name = "configuration",
    srcs = [
        "configuration.go",
    ],
    deps = [
        "//consumers/jira_c/types:config",
        "//third_party/go:yaml.v2",
    ],
)

go_test(
    name = "configuration_test",
    srcs = [
        "configuration_test.go",
    ],
    deps = [
        ":configuration",
        "//third_party/go:stretchr_testify",
        "//third_party/go:yaml.v2",
        "//consumers/jira_c/types:config",
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
