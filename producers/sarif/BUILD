subinclude("//third_party/defs:docker")

# this producer covers any tool that supports the SARIF format https://www.oasis-open.org/committees/tc_home.php?wg_abbrev=sarif
go_binary(
    name = "sarif",
    srcs = [
        "main.go",
    ],
    deps = [
        "//api/proto:v1",
        "//producers",
        "//third_party/go:go-sarif",
    ],
)

go_test(
    name = "sarif_test",
    srcs = [
        "main.go",
        "main_test.go",
    ],
    deps = [
        "//api/proto:v1",
        "//producers",
        "//third_party/go:go-sarif",
        "//third_party/go:stretchr_testify",
    ],
)

docker_image(
    name = "dracon-producer-sarif",
    srcs = [
        ":sarif",
    ],
    base_image = "//build/docker:dracon-base-go",
    image = "dracon-producer-sarif",
)
