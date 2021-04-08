subinclude("//build/defs:kustomize")

kustomized_config(
    name = "tektoncd-dashboard",
    srcs = [
        "//third_party/k8s:tektoncd_dashboard",
        "ingress.yaml",
        "kustomization.yaml",
    ],
    visibility = ["//scripts/development/..."],
)
