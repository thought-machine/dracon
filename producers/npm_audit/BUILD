subinclude("@third_party/subrepos/pleasings//docker")

go_binary(
    name = "npm_audit",
    srcs = [
        "main.go",
    ],
    deps = [
        "//api/proto:v1",
        "//producers",
        "//producers/npm_audit/types:npmaudit-issue"
    ],
)

go_test(
    name = "npm_audit_test",
    srcs = [
        "main_test.go",
        "main.go"
    ],
    deps = [
        "//api/proto:v1",
        "//producers",
        "//third_party/go:stretchr_testify",
        "//producers/npm_audit/types:npmaudit-issue",
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
