KUSTOMIZE_VERSION = "4.2.0"
remote_file(
    name = "kustomize",
    url = f"https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2Fv{KUSTOMIZE_VERSION}/kustomize_v{KUSTOMIZE_VERSION}_{CONFIG.OS}_{CONFIG.ARCH}.tar.gz",
    binary = True,
    extract = True,
    visibility = ["PUBLIC"],
    hashes = [
        "220dd03dcda8e45dc50e4e42b2d71882cbc4c05e0ed863513e67930ecad939eb", # linux-amd64
    ],
)
