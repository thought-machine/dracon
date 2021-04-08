subinclude("//third_party/defs:docker")

filegroup(
    name = "npm_audit",
    srcs = [
        "index.js",
        "package.json",
        "package-lock.json",
    ],
)

docker_image(
    name = "npm-audit",
    srcs = [
        ":npm_audit",
    ],
    dockerfile = "Dockerfile-tool-npm-audit",
    image = "dracon-tool-npm-audit",
)
