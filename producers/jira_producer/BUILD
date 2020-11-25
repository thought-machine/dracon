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
