#!/bin/bash
set -Eeuo pipefail

clear_remainder="\033[0K"

util::info() {
    printf "$(ansi::resetColor)$(ansi::magentaIntense)💡 %s$(ansi::resetColor)\n" "$@"
}

util::infor() {
    printf "$(ansi::resetColor)$(ansi::magentaIntense)💡 %s$(ansi::resetColor)" "$@"
}

util::rinfor() {
    printf "\r$(ansi::resetColor)$(ansi::magentaIntense)💡 %s$(ansi::resetColor)${clear_remainder}" "$@"
}

util::warn() {
  printf "$(ansi::resetColor)$(ansi::yellowIntense)⚠️  %s$(ansi::resetColor)\n" "$@"
}

util::error() {
  printf "$(ansi::resetColor)$(ansi::bold)$(ansi::redIntense)❌ %s$(ansi::resetColor)\n" "$@"
}

util::rerror() {
  printf "\r$(ansi::resetColor)$(ansi::bold)$(ansi::redIntense)❌ %s$(ansi::resetColor)${clear_remainder}\n" "$@"
}

util::success() {
  printf "$(ansi::resetColor)$(ansi::greenIntense)✅ %s$(ansi::resetColor)\n" "$@"
}

util::rsuccess() {
  printf "\r$(ansi::resetColor)$(ansi::greenIntense)✅ %s$(ansi::resetColor)${clear_remainder}\n" "$@"
}

util::retry() {
  "${@}" || sleep 1; "${@}" || sleep 5; "${@}"
}

util::prompt() {
  prompt=$(printf "$(ansi::bold)$()❔ %s [y/N]$(ansi::resetColor)\n" "$@")
  read -p "${prompt}" yn
  case $yn in
      [Yy]* ) ;;
      * ) util::error "Did not receive happy input, exiting."; exit 1;;
  esac
}

util::prompt_skip() {
  prompt=$(printf "$(ansi::bold)$()❔ %s [y/N]$(ansi::resetColor)\n" "$@")
  read -p "${prompt}" yn
  case $yn in
      [Yy]* ) return 0;;
      * ) util::warn "Did not receive happy input, skipping."; return 1;;
  esac
}

util::waitForRollout() {
  local k8s_yaml kind namespace resource limit attempts
  k8s_yaml="$1"
  limit=100 # 300 seconds

  kind_namespace_resources=($("$YQ_BIN" e -N '[. = .kind + "/" + .metadata.namespace + "/" + .metadata.name]' "$k8s_yaml" | sed 's/\/\//\/default\//g'))

  for knr in "${kind_namespace_resources[@]}"; do
    kind="$(echo "$knr" | cut -d/ -f1)"
    namespace="$(echo "$knr" | cut -d/ -f2)"
    resource="$(echo "$knr" | cut -d/ -f3)"
    attempts=0
    if [[ "$kind" =~ ^(Deployment|Statefulset)$ ]]; then
      util::rinfor "waiting for deployment ${namespace}/${resource}"
      rollout_status_cmd="kubectl -n ${namespace} rollout status ${kind}/${resource}"
      
      until $rollout_status_cmd > /dev/null || [ $attempts -eq $limit ]; do
        attempts=$((attempts + 1))
        echo "foo"
        kubectl -n "${namespace}" logs "${kind,,}/${resource}"

        sleep 3
      done
    fi
    if [ $attempts -eq $limit ]; then
      util::rerror "Deployment '${namespace}/${resource}' did not roll out $($rollout_status_cmd)"

      kubectl -n "${namespace}" describe "${kind}" "${resource}"
      kubectl -n "${namespace}" logs "${kind,,}/${resource}"
      exit 1
    fi
  done  
}
