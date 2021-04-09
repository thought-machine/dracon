KUBERNETES_INGRESSNGINX_VERSION = "3.25.0"

remote_file(
    name = "kubernetes_ingressnginx",
    url = f"https://raw.githubusercontent.com/kubernetes/ingress-nginx/helm-chart-{KUBERNETES_INGRESSNGINX_VERSION}/deploy/static/provider/kind/deploy.yaml",
    visibility = ["//scripts/development/kind/..."],
)

JETSTACK_CERTMANAGER_VERSION = "1.1.0"

remote_file(
    name = "jetstack_certmanager",
    url = f"https://github.com/jetstack/cert-manager/releases/download/v{JETSTACK_CERTMANAGER_VERSION}/cert-manager.yaml",
    visibility = ["//scripts/development/kind/..."],
)

TEKTONCD_PIPELINE_VERSION = "0.22.0"

remote_file(
    name = "tektoncd_pipeline",
    url = f"https://github.com/tektoncd/pipeline/releases/download/v{TEKTONCD_PIPELINE_VERSION}/release.yaml",
    visibility = ["//scripts/development/..."],
)

TEKTONCD_DASHBOARD_VERSION = "0.15.0"

remote_file(
    name = "tektoncd_dashboard",
    url = f"https://github.com/tektoncd/dashboard/releases/download/v{TEKTONCD_DASHBOARD_VERSION}/tekton-dashboard-release-readonly.yaml",
    visibility = ["//scripts/development/..."],
)
