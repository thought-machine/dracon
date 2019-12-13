package template

import (
	"encoding/json"
	"testing"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/stretchr/testify/assert"
)

func TestArrayPatch(t *testing.T) {
	startPatchBytes := []byte(`
[
	{
		"op": "add",
		"path": "/spec/steps/*/volumeMounts/-",
		"value": {
			"mountPath": "/dracon",
			"name": "dracon-ws"
		}
	}
]
`)
	startPatch, err := jsonpatch.DecodePatch(startPatchBytes)
	assert.Nil(t, err)

	targetJSON := []byte(`
{
	"spec": {
		"steps": [
			{
				"volumeMounts": []
			},
			{
				"volumeMounts": []
			}
		]
	}
}`)

	resPatch, err := patchArrayGlob(startPatch[0], targetJSON)
	assert.Nil(t, err)

	resPatchJSON, err := json.Marshal(resPatch)
	assert.Nil(t, err)

	expectedPatchJSON := []byte(`
[
	{
		"op": "add",
		"path":"/spec/steps/0/volumeMounts/-",
		"value": {
			"mountPath": "/dracon",
			"name": "dracon-ws"
		}
	},
	{
		"op": "add",
		"path":"/spec/steps/1/volumeMounts/-",
		"value": {
			"mountPath": "/dracon",
			"name": "dracon-ws"
		}
	}
]
`)

	assert.JSONEq(t, string(expectedPatchJSON), string(resPatchJSON))
}
