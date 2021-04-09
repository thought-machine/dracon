subinclude("//build/defs:kustomize")

kustomized_config(
    name = "tektoncd-dashboard",
    srcs = [
        "ingress.yaml",
        "kustomization.yaml",
        "//third_party/k8s:tektoncd_dashboard",
    ],
    visibility = ["//scripts/development/..."],
)
