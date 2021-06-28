KUSTOMIZE_VERSION = "4.1.3"
remote_file(
    name = "kustomize",
    url = f"https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2Fv{KUSTOMIZE_VERSION}/kustomize_v{KUSTOMIZE_VERSION}_{CONFIG.OS}_{CONFIG.ARCH}.tar.gz",
    binary = True,
    extract = True,
    hashes = [
        "f028cd2b675d215572d54634311777aa475eb5612fb8a70d84b957c4a27a861e", # linux-amd64
    ],
)
