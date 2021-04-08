#!/bin/bash
# This script deploys the given plz target that outputs k8s resources to the given kubernetes context.
set -Eeuo pipefail

DEFINE_boolean 'wait' false 'wait for pods to be ready' 'w'
FLAGS "$@" || exit $?
eval set -- "${FLAGS_ARGV}"

target="$1"

if [ -z "${target}" ]; then
  util::error "missing target"
  exit 1
fi

kubernetes_context=$(kubectl config current-context)
if [ "${kubernetes_context}" != "kind-dracon" ]; then
  util::prompt "Are you sure you would like to deploy ${target} to '${kubernetes_context}'?"
fi

kind_cluster_name="${kubernetes_context//kind-/}"

util::load_development_images_into_kind "${kind_cluster_name}" "${target}"

util::infor "configuring CA"

ca_crt_b64=$(base64 -w0 < "${ROOT_CERTIFICATES}/ca.crt")
ca_key_b64=$(base64 -w0 < "${ROOT_CERTIFICATES}/ca.key")
ca_apply_out=$(cat <<EOF | kubectl --context "${kubernetes_context}" apply -f -
apiVersion: v1
kind: Secret
type: kubernetes.io/tls
metadata:
  name: localhost-issuer-ca-key
  namespace: cert-manager
data:
  tls.crt: ${ca_crt_b64}
  tls.key: ${ca_key_b64}
EOF
)

util::retry kubectl --context "${kubernetes_context}" apply -f "${CERTMANAGER_CONFIG}" > /dev/null

if [[ "${ca_apply_out}" == *"configured"* ]]; then
  util::rinfor "renewing certs"
  # renew certs
  kubectl --context "${kubernetes_context}" --namespace dracon get certs --no-headers=true | awk '{print $1}' | xargs -n 1 kubectl --context "${kubernetes_context}" --namespace dracon patch certificate --patch '
  - op: replace
    path: /spec/renewBefore
    value: 2159h59m59s
  ' --type=json

  kubectl --context "${kubernetes_context}" --namespace dracon get certs --no-headers=true | awk '{print $1}' | xargs -n 1 kubectl --context "${kubernetes_context}" --namespace dracon patch certificate --patch '
  - op: remove
    path: /spec/renewBefore
  ' --type=json
fi
util::rinfor "configured CA"

util::rinfor "building Kubernetes resource"
./pleasew build "${target}" > /dev/null
util::rinfor "built Kubernetes resource"

k8s_out=($(./pleasew query output "${target}"))

util::rinfor "deploying ${target} resource ${kubernetes_context}"
for k8s_yaml in "${k8s_out[@]}"; do
    kubectl --context "${kubernetes_context}" apply -f "${k8s_yaml}" > /dev/null
done

if [ ${FLAGS_wait} -eq ${FLAGS_TRUE} ]; then
  namespace="default"
  for k8s_yaml in "${k8s_out[@]}"; do
      util::waitForRollout "${k8s_yaml}"
  done
fi

util::rsuccess "deployed ${target} resources to ${kubernetes_context}"
