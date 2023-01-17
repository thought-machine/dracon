go_library(
    name = "template",
    srcs = [
        "array-patch.go",
        "load.go",
        "patch.go",
        "pipeline.go",
        "pipeline-resource.go",
        "template.go",
    ],
    resources = [":resources"],
    visibility = ["PUBLIC"],
    deps = [
        "//third_party/go:apimachinery",
        "//third_party/go:evanphx_json-patch",
        "//third_party/go:ghodss_yaml",
        "//third_party/go:speps_go-hashids",
        "//third_party/go:yaml_v3",
    ],
)

go_test(
    name = "template_test",
    srcs = [
        "load_test.go",
    ],
    resources = [":resources"],
    deps = [
        ":template",
        "//third_party/go:heredoc",
        "//third_party/go:stretchr_testify",
    ],
)

# This is a workaround for embedding files that are
# not in the same pkg dir.
filegroup(
    name = "resources",
    srcs = ["//resources/patches"],
)
