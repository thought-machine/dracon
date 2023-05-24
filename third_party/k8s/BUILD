KUBERNETES_INGRESSNGINX_VERSION = "3.25.0"

remote_file(
    name = "kubernetes_ingressnginx",
    hashes = ["5cea6ce0f7e285f88847852aab05902406a63bfe4aef297831ea0207111fae44"],
    url = f"https://raw.githubusercontent.com/kubernetes/ingress-nginx/helm-chart-{KUBERNETES_INGRESSNGINX_VERSION}/deploy/static/provider/kind/deploy.yaml",
    visibility = ["//scripts/development/kind/..."],
)

JETSTACK_CERTMANAGER_VERSION = "1.1.0"

remote_file(
    name = "jetstack_certmanager",
    hashes = ["651857a32b0be92f3c5c274729eea9c34bd457726bd3f560b3f68399659d1cfd"],
    url = f"https://github.com/jetstack/cert-manager/releases/download/v{JETSTACK_CERTMANAGER_VERSION}/cert-manager.yaml",
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
    hashes = ["48d26b523438eb35bb622270b6ae7853f9238f0703f5961297199edd001d0529"],
    out = "tekton-dashboard-release-readonly.yaml",
    url = f"https://github.com/tektoncd/dashboard/releases/download/v{TEKTONCD_DASHBOARD_VERSION}/release.yaml",
    visibility = ["//scripts/development/..."],
)
