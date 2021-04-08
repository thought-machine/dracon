subinclude("//third_party/defs:docker")

go_binary(
    name = "sync_tickets",
    srcs = [
        "main.go",
    ],
    deps = [
        "//common/jira/config",
        "//common/jira/jira",
        "//pkg/enrichment/db",
        "//producers/jira_producer/sync",
        "//third_party/go:go-jira",
        "//third_party/go:spf13_cobra",
        "//third_party/go:spf13_viper",
    ],
)

docker_image(
    name = "dracon-producer-jira",
    srcs = [
        ":sync_tickets",
    ],
    base_image = "//build/docker:dracon-base-go",
    image = "dracon-producer-jira",
)
