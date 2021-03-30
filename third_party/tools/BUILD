KIND_VERSION = "0.10.0"

remote_file(
    name = "kind",
    binary = True,
    hashes = [
        "74767776488508d847b0bb941212c1cb76ace90d9439f4dee256d8a04f1309c6",  # linux-amd64
    ],
    url = f"https://github.com/kubernetes-sigs/kind/releases/download/v{KIND_VERSION}/kind-{CONFIG.OS}-{CONFIG.ARCH}",
    visibility = ["//scripts/..."],
)

remote_file(
    name = "yq",
    binary = True,
    hashes = [
        "0e105edbb0ebc7c4115c610168f1d6b0ff3ceb38043dac92e18fa0698130d69f", # linux_amd64
    ],
    url = f"https://github.com/mikefarah/yq/releases/download/v4.6.3/yq_{CONFIG.OS}_{CONFIG.ARCH}.tar.gz",
    extract = True,
    visibility = ["//scripts/..."],
)
