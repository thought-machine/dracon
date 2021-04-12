#!/bin/bash
# This script creates a KinD cluster for dracon.
set -Eeuo pipefail

DEFINE_string 'kind_cluster' 'dracon' 'Kind cluster to use' 'c'
DEFINE_string 'kubernetes_version' '1.17' 'Kubernetes version to use' 'k'
FLAGS "$@" || exit $?
eval set -- "${FLAGS_ARGV}"

kubernetes_version="${FLAGS_kubernetes_version//./_}"

declare -A KUBERNETES_VERSIONS
KUBERNETES_VERSIONS=(
  ["1_16"]="kindest/node:v1.16.15@sha256:c10a63a5bda231c0a379bf91aebf8ad3c79146daca59db816fb963f731852a99"
  ["1_17"]="kindest/node:v1.17.17@sha256:7b6369d27eee99c7a85c48ffd60e11412dc3f373658bc59b7f4d530b7056823e"
  ["1_18"]="kindest/node:v1.18.15@sha256:5c1b980c4d0e0e8e7eb9f36f7df525d079a96169c8a8f20d8bd108c0d0889cc4"
)

kubernetes_context="kind-${FLAGS_kind_cluster}"

if ! $KIND_BIN get clusters | grep "${FLAGS_kind_cluster}" 2>&1 > /dev/null; then
  util::info "Creating KinD cluster ${FLAGS_kind_cluster} (${FLAGS_kubernetes_version})"
  sed -i "s#___KIND_IMAGE___#${KUBERNETES_VERSIONS[$kubernetes_version]}#" "${KIND_CONFIG}"
  $KIND_BIN create cluster --name dracon --config "${KIND_CONFIG}"
  util::success "Created KinD cluster"
  kubectl config use-context "${kubernetes_context}"
else
  util::success "KinD cluster already exists"
fi

kubectl="kubectl --context ${kubernetes_context}"

# kubernetes/ingress-nginx
util::infor "Installing kubernetes/ingress-nginx"
$kubectl apply -f "${KUBERNETES_INGRESSNGINX_INSTALL}" > /dev/null
$kubectl apply -f "${KUBERNETES_INGRESSNGINX_CONFIG}" > /dev/null

util::rinfor "waiting for kubernetes/ingress-nginx pods"
util::waitForRollout "${KUBERNETES_INGRESSNGINX_INSTALL}"
util::rsuccess "Installed kubernetes/ingress-nginx"
  
# jetstack/cert-manager
util::infor "Installing jetstack/cert-manager"
$kubectl apply -f "${JETSTACK_CERTMANAGER_INSTALL}" > /dev/null
util::rinfor "waiting for jetstack/cert-manager pods"
util::waitForRollout "${JETSTACK_CERTMANAGER_INSTALL}"
util::rsuccess "Installed jetstack/cert-manager"

# tektoncd/pipeline
util::infor "Installing tektoncd/pipeline"
$kubectl apply -f "${TEKTONCD_PIPELINE_INSTALL}" > /dev/null
util::waitForRollout "${TEKTONCD_PIPELINE_INSTALL}"
util::rsuccess "Installed tektoncd/pipeline"

# tektoncd/dashboard
util::infor "Installing tektoncd/dashboard"
$kubectl apply -f "${TEKTONCD_DASHBOARD_INSTALL}" > /dev/null
util::waitForRollout "${TEKTONCD_DASHBOARD_INSTALL}"
util::rsuccess "Installed tektoncd/dashboard"


# Finish
if [ "$(kubectl config current-context)" != "${kubernetes_context}" ]; then
  util::warn "Current kubectl context is not ${kubernetes_context}. You can set this with:
  $ kubectl config use-context \"${kubernetes_context}\"
"
fi
