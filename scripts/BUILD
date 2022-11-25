genrule(
    name = "default_docker_repo",
    outs = ["default_docker_repo"],
    cmd = f'echo -n "{CONFIG.DEFAULT_DOCKER_REPO}" > $OUTS',
)

for script in glob(["*.sh"]):
    sh_binary(
        name = basename(splitext(script)[0]),
        main = script,
        data = [
            ":default_docker_repo",
        ],
        deps = [
            "//third_party/sh:shflags",
        ],
    )
