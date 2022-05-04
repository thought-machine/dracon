KIND_VERSION = "0.12.0"

remote_file(
    name = "kind",
    binary = True,
    hashes = [
        "b80624c14c807490c0944d21fdc9c3455d6cc904fad486fe236f2187ecaa5789",  # linux-amd64
    ],
    url = f"https://github.com/kubernetes-sigs/kind/releases/download/v{KIND_VERSION}/kind-{CONFIG.OS}-{CONFIG.ARCH}",
    visibility = ["//scripts/..."],
)

remote_file(
    name = "yq",
    binary = True,
    extract = True,
    hashes = [
        "0e105edbb0ebc7c4115c610168f1d6b0ff3ceb38043dac92e18fa0698130d69f",  # linux_amd64
    ],
    url = f"https://github.com/mikefarah/yq/releases/download/v4.6.3/yq_{CONFIG.OS}_{CONFIG.ARCH}.tar.gz",
    visibility = ["//scripts/..."],
)

KUSTOMIZE_VERSION = "v3.8.7"

remote_file(
    name = "kustomize",
    binary = True,
    extract = True,
    hashes = [
        "4a3372d7bfdffe2eaf729e77f88bc94ce37dc84de55616bfe90aac089bf6fd02",  # linux-amd64
    ],
    url = f"https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2F{KUSTOMIZE_VERSION}/kustomize_{KUSTOMIZE_VERSION}_{CONFIG.OS}_{CONFIG.ARCH}.tar.gz",
    visibility = [
        "//examples/...",
        "//scripts/development/k8s/...",
    ],
)
