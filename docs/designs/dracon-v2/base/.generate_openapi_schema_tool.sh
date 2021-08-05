#!/bin/bash

set -euo pipefail

JQ_BIN="$(dirname $0)/docs/designs/dracon-v2/third_party/binary/jq-linux64"

elementIn () {
  local e match="$1"
  shift
  for e; do [[ "$e" == "$match" ]] && return 0; done
  return 1
}

# only use .definitions
cp $SRCS $OUTS

# add x-kubernetes-group-version-kind to all objects
crds=($($JQ_BIN -r '.definitions | keys | .[]' $OUTS))
for crd in "${crds[@]}"; do
    gvk_crd="${crd//v1alpha1/tekton.dev\/v1alpha1}"
    gvk_crd="${gvk_crd//v1beta1/tekton.dev\/v1beta1}"
    kind="$(echo $gvk_crd | rev | cut -d\. -f1 | rev)"
    version="$(echo $gvk_crd | rev | cut -d\. -f2 | rev | cut -d\/ -f2)"
    group="$(echo $gvk_crd | cut -d\/ -f1)"
    if [ -n "$group" ]; then
        $JQ_BIN \
            --arg group "$group" \
            --arg version "$version" \
            --arg kind "$kind" \
            --arg crd "$crd" \
        '.definitions[$crd] += { "x-kubernetes-group-version-kind": [{ "group": $group, "version": $version, "kind": $kind }] }' $OUTS > $OUTS.new
        mv $OUTS.new $OUTS
    fi
done

# remove v1alpha1 crds
for crd in "${crds[@]}"; do
    if [[ $crd == v1alpha1* ]]; then
        echo "removing $crd"
        $JQ_BIN --arg crd "$crd" 'del(.definitions[$crd])' $OUTS > $OUTS.new
        mv $OUTS.new $OUTS
    fi
done
crds=($($JQ_BIN -r '.definitions | keys | .[]' $OUTS))

# strip \$refs which don't reference in the same schema
all_refs=($($JQ_BIN -r 'paths | join("][")' $OUTS | grep "\$ref$" | while read line; do echo "[${line}]"; done | sed -e 's|\[|\["|g' | sed -e 's|\]|"\]|g'))
for ref in "${all_refs[@]}"; do
    ref_value=$($JQ_BIN -r ".$ref" $OUTS | sed 's|#/definitions/||g')
    if ! elementIn "${ref_value}" "${crds[@]}"; then
        echo "removing: ${ref} (${ref_value} does not exist)"
        $JQ_BIN "del(.$ref)" $OUTS > $OUTS.new
        mv $OUTS.new $OUTS
    fi
done

# add x-kubernetes-patch-merge-key and x-kubernetes-patch-strategy
# allow the Pipeline.spec.tasks list to be merged
$JQ_BIN '.definitions["v1beta1.PipelineSpec"].properties.tasks += { "x-kubernetes-patch-strategy": "merge" }' $OUTS > $OUTS.new && mv $OUTS.new $OUTS
$JQ_BIN '.definitions["v1beta1.PipelineSpec"].properties.tasks += { "x-kubernetes-patch-merge-key": "name" }' $OUTS > $OUTS.new && mv $OUTS.new $OUTS

# allow the Pipeline.spec.tasks[].params to be merged
$JQ_BIN '.definitions["v1beta1.PipelineTask"].properties.params += { "x-kubernetes-list-type": "map" }' $OUTS > $OUTS.new && mv $OUTS.new $OUTS
$JQ_BIN '.definitions["v1beta1.PipelineTask"].properties.params += { "x-kubernetes-patch-strategy": "merge" }' $OUTS > $OUTS.new && mv $OUTS.new $OUTS
$JQ_BIN '.definitions["v1beta1.PipelineTask"].properties.params += { "x-kubernetes-patch-merge-key": "name" }' $OUTS > $OUTS.new && mv $OUTS.new $OUTS

# butcher the Pipeline.spec.tasks[].params[].value so that they can be merged additively
$JQ_BIN '.definitions["v1beta1.Param"].properties.value += { "type": "array", "items": { "type": "string" } }' $OUTS > $OUTS.new && mv $OUTS.new $OUTS
$JQ_BIN 'del(.definitions["v1beta1.Param"].properties.value["$ref"])' $OUTS > $OUTS.new && mv $OUTS.new $OUTS
$JQ_BIN 'del(.definitions["v1beta1.Param"].properties.value.default)' $OUTS > $OUTS.new && mv $OUTS.new $OUTS
$JQ_BIN '.definitions["v1beta1.Param"].properties.value += { "x-kubernetes-patch-strategy": "merge" }' $OUTS > $OUTS.new && mv $OUTS.new $OUTS
$JQ_BIN '.definitions["v1beta1.Param"].properties.value += { "x-kubernetes-list-type": "map" }' $OUTS > $OUTS.new && mv $OUTS.new $OUTS


# allow the Task.spec.steps list to be merged
$JQ_BIN '.definitions["v1beta1.TaskSpec"].properties.steps += { "x-kubernetes-patch-strategy": "merge" }' $OUTS > $OUTS.new && mv $OUTS.new $OUTS
$JQ_BIN '.definitions["v1beta1.TaskSpec"].properties.steps += { "x-kubernetes-patch-merge-key": "name" }' $OUTS > $OUTS.new && mv $OUTS.new $OUTS
