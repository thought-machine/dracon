subinclude("//third_party/defs:docker")

go_binary(
    name = "npm_audit",
    srcs = [
        "main.go",
    ],
    deps = [
        "//api/proto:v1",
        "//producers",
        "//producers/npm_audit/types",
        "//producers/npm_audit/types/npm_full_audit",
        "//producers/npm_audit/types/npm_quick_audit",
    ],
)

docker_image(
    name = "dracon-producer-npmaudit",
    srcs = [
        ":npm_audit",
    ],
    base_image = "//build/docker:dracon-base-go",
    image = "dracon-producer-npm-audit",
)
