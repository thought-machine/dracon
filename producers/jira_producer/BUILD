subinclude("//third_party/defs:docker")

go_binary(
    name = "sync_tickets",
    srcs = [
        "main.go",
    ],
    deps = [
                "//producers/jira_producer/sync",
        "//third_party/go:spf13_viper",
        "//third_party/go:spf13_cobra",
        "//common/jira/jira:jira",
        "//pkg/enrichment/db",
        "//third_party/go:go-jira",
        "//common/jira/config"
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
