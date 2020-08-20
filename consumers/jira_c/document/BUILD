go_library(
    name = "document",
    srcs = [
        "document.go",
    ],
    deps = [
        ":types",
        "//third_party/go:protobuf",
        "//api/proto:v1",
    ],
    visibility = [
        "//consumers/jira_c/utils:utils"
    ],
)

go_library(
    name = "types",
    srcs = [
        "types.go",
    ],
    visibility = [
        "//consumers/jira_c/..."
    ],
)