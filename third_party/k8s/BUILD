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

TEKTONCD_PIPELINE_VERSION = "0.19.0"

remote_file(
    name = "tektoncd_pipeline",
    hashes = ["a6a9049595137ec95cd50f8c0767a01cf54e941f7b2381710596f1bc4f8db29b"],
    url = f"https://github.com/tektoncd/pipeline/releases/download/v{TEKTONCD_PIPELINE_VERSION}/release.yaml",
    visibility = ["//scripts/development/..."],
)

TEKTONCD_DASHBOARD_VERSION = "0.14.0"

remote_file(
    name = "tektoncd_dashboard",
    hashes = ["938edf9e16fdc91585a944b8f50182638fd5b653ac1311bfac73be768e2a730c"],
    url = f"https://github.com/tektoncd/dashboard/releases/download/v{TEKTONCD_DASHBOARD_VERSION}/tekton-dashboard-release-readonly.yaml",
    visibility = ["//scripts/development/..."],
)
