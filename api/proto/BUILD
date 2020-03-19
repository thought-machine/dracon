
proto_library(
    name = "config",
    srcs = ["config.proto"],
    languages = ["py"],
    visibility=["PUBLIC"],
)


proto_library(
    name = "engine",
    srcs = ["engine.proto"],
    languages = ["py"],
    visibility=["PUBLIC"],
    deps = [
            ":issue",
            ":config",
            ]
)

proto_library(
    name = "issue",
    srcs = ["issue.proto"],
    languages = ["py"],
    visibility=["PUBLIC"],
)

