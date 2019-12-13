go_library(
    name = "cmd",
    srcs = [
        "root.go",
        "run.go",
        "setup.go",
        "version.go",
    ],
    visibility = ["PUBLIC"],
    deps = [
        "//pkg/kubernetes",
        "//pkg/template",
        "//pkg/version",
        "//third_party/go:mitchellh_go-homedir",
        "//third_party/go:spf13_cobra",
        "//third_party/go:spf13_viper",
    ],
)
