KUBERNETES_INGRESSNGINX_VERSION = "4.5.2"

remote_file(
    name = "kubernetes_ingressnginx",
    hashes = ["c36d92bf32f76fac22aa9324f15d1019a7380411f07261f903a713fc85737fbb"],
    url = f"https://raw.githubusercontent.com/kubernetes/ingress-nginx/helm-chart-{KUBERNETES_INGRESSNGINX_VERSION}/deploy/static/provider/kind/deploy.yaml",
    visibility = ["//scripts/development/kind/..."],
)

CERTMANAGER_CERTMANAGER_VERSION = "1.12.0"

remote_file(
    name = "certmanager_certmanager",
    hashes = ["fac060a5a953d2ebe2f38e3ded79ebbee3bebe9d26078478618d81bebd9c6385"],
    url = f"https://github.com/cert-manager/cert-manager/releases/download/v{CERTMANAGER_CERTMANAGER_VERSION}/cert-manager.yaml",
    visibility = ["//scripts/development/kind/..."],
)

TEKTONCD_PIPELINE_VERSION = "0.45.0"

remote_file(
    name = "tektoncd_pipeline",
    hashes = ["b4c6a5d7f8e8692095751ab98eb272f84cb5ace8e1084728b19a0bc1263ff614"],
    url = f"https://github.com/tektoncd/pipeline/releases/download/v{TEKTONCD_PIPELINE_VERSION}/release.yaml",
    visibility = ["//scripts/development/..."],
)

TEKTONCD_DASHBOARD_VERSION = "0.34.0"

remote_file(
    name = "tektoncd_dashboard",
    out = "tekton-dashboard-release-readonly.yaml",
    hashes = ["48d26b523438eb35bb622270b6ae7853f9238f0703f5961297199edd001d0529"],
    url = f"https://github.com/tektoncd/dashboard/releases/download/v{TEKTONCD_DASHBOARD_VERSION}/release.yaml",
    visibility = ["//scripts/development/..."],
)
