#!/bin/bash
# This script creates a KinD cluster for dracon.
set -Eeuo pipefail

DEFINE_string 'kind_cluster' 'dracon' 'Kind cluster to use' 'c'
FLAGS "$@" || exit $?
eval set -- "${FLAGS_ARGV}"

kubernetes_context="kind-${FLAGS_kind_cluster}"

if ! $KIND_BIN get clusters | grep "${FLAGS_kind_cluster}" 2>&1 > /dev/null; then
  util::info "Creating KinD cluster ${FLAGS_kind_cluster}"
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


# Finish
if [ "$(kubectl config current-context)" != "${kubernetes_context}" ]; then
  util::warn "Current kubectl context is not ${kubernetes_context}. You can set this with:
  $ kubectl config use-context \"${kubernetes_context}\"
"
fi
