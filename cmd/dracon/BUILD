go_binary(
    name = "dracon",
    srcs = [
        "main.go",
    ],
    definitions = {
        "github.com/thought-machine/dracon/pkg/version.BuildVersion": "${SCM_REVISION}",
    },
    stamp = True,
    static = True,
    deps = [
        "//cmd/dracon/cmd",
        "//third_party/go:spf13_cobra",
    ],
)
