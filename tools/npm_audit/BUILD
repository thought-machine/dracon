subinclude("@third_party/subrepos/pleasings//docker")

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
