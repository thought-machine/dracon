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
    visibility = ["PUBLIC"],
    deps = [
        "//pkg/template/patches",
        "//third_party/go:apimachinery",
        "//third_party/go:evanphx_json-patch",
        "//third_party/go:ghodss_yaml",
        "//third_party/go:rakyll_statik",
        "//third_party/go:speps_go-hashids",
    ],
)
