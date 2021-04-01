sh_cmd(
    name = "setup",
    srcs = ["setup.sh"],
    cmd = """
source "$(out_location //scripts/util)"
source "$(out_location //third_party/sh:shflags)"

KIND_BIN="$(out_exe //third_party/tools:kind)"
KIND_CONFIG="$(out_location :_configuration)"

KUBERNETES_INGRESSNGINX_INSTALL="$(out_location //third_party/k8s:kubernetes_ingressnginx)"
KUBERNETES_INGRESSNGINX_CONFIG="$(out_location //scripts/development/kind/ingress:ingress-nginx)"

JETSTACK_CERTMANAGER_INSTALL="$(out_location //third_party/k8s:jetstack_certmanager)"

TEKTONCD_PIPELINE_INSTALL="$(out_location //third_party/k8s:tektoncd_pipeline)"
TEKTONCD_DASHBOARD_INSTALL="$(out_location //scripts/development/k8s/tektoncd-dashboard)"

source $SRCS
    """,
    data = [
        "//scripts/util",
        "//third_party/tools:kind",
        "//third_party/sh:shflags",
        "//third_party/k8s:kubernetes_ingressnginx",
        "//third_party/k8s:jetstack_certmanager",
        "//scripts/development/kind/ingress:ingress-nginx",
        "//third_party/k8s:tektoncd_pipeline",
        "//scripts/development/k8s/tektoncd-dashboard",
        ":_configuration",
    ],
    shell = "/bin/bash",
)

sh_cmd(
    name = "delete",
    cmd = """
source $(out_location //scripts/util)

KIND_BIN="$(out_location //third_party/tools:kind)"

util::info "Deleting KinD cluster dracon"
\\\\$KIND_BIN delete cluster --name dracon

util::success "Deleted KinD cluster"
    """,
    data = [
        "//scripts/util",
        "//third_party/tools:kind",
    ],
    shell = "/bin/bash",
)

sh_cmd(
    name = "deploy",
    srcs = ["deploy.sh"],
    cmd = """
source "$(out_location :_util)"
source "$(out_location //third_party/sh:shflags)"

ROOT_CERTIFICATES="$(out_location :_root_certificates)"
CERTMANAGER_CONFIG="$(out_location //scripts/development/kind/ingress:cert-manager)"

KIND_BIN="$(out_location //third_party/tools:kind)"
YQ_BIN="$(out_location //third_party/tools:yq)"

source $SRCS
    """,
    data = [
        ":_util",
        "//third_party/tools:kind",
        "//third_party/sh:shflags",
        ":_root_certificates",
        "//third_party/tools:yq",
        "//scripts/development/kind/ingress:cert-manager",
    ],
    shell = "/bin/bash",
)

sh_cmd(
    name = "dracon",
    srcs = ["dracon.sh"],
    cmd = """
source "$(out_location :_util)"

DRACON_BIN=$(out_exe //cmd/dracon)
KIND_BIN=$(out_exe //third_party/tools:kind)
source $SRCS
    """,
    data = [
        "//cmd/dracon",
        "//third_party/tools:kind",
        ":_util",
    ],
    shell = "/bin/bash",
)

sh_cmd(
    name = "_util",
    srcs = ["util.sh"],
    cmd = """
source "$(out_location //scripts/util)"

source $SRCS
    """,
    data = [
        "//scripts/util",
    ],
    shell = "/bin/bash",
)


filegroup(
    name = "_configuration",
    srcs = ["configuration.yaml"],
)

genrule(
    name = "_root_certificates",
    outs = ["_root_certificates"],
    cmd = """
openssl genrsa -out ca.key 2048
openssl req -x509 -new -nodes \\
    -key ca.key \\
    -out ca.crt \\
    -subj "/CN=dracon.localhost" \\
    -days 14 \\
    -reqexts v3_req -extensions v3_ca
mkdir $OUTS
mv ca.key $OUTS
mv ca.crt $OUTS
    """,
)
