package template

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	jsonpatch "github.com/evanphx/json-patch"
)

func patchArrayGlob(op jsonpatch.Operation, targetJSON []byte) (jsonpatch.Patch, error) {
	path, err := op.Path()
	if err != nil || !strings.Contains(path, `/*/`) {
		return jsonpatch.Patch{op}, nil
	}
	return resolveArrayGlobOps(op, targetJSON)
}

func resolveArrayGlobOps(op jsonpatch.Operation, targetJSON []byte) (jsonpatch.Patch, error) {
	resOps := jsonpatch.Patch{}
	path, err := op.Path()
	if err != nil {
		return nil, fmt.Errorf("could not determine operation path: %w", err)
	}

	getLengthOfPath := func() (int, error) {
		var objMap map[string]*json.RawMessage
		err := json.Unmarshal(targetJSON, &objMap)
		if err != nil {
			return -1, fmt.Errorf("could parse target JSON: %w", err)
		}

		pathParts := strings.Split(path, "/")
		pathParts = pathParts[1:]
		for i, key := range pathParts {
			if rawJSON, ok := objMap[key]; ok {
				if pathParts[i+1] == "*" {
					var objArr []*json.RawMessage
					json.Unmarshal(*rawJSON, &objArr)
					return len(objArr), nil
				}
				err := json.Unmarshal(*rawJSON, &objMap)
				if err != nil {
					return -1, fmt.Errorf("could parse target JSON: %w", err)
				}
			}
		}
		return 0, nil
	}

	resolvedLength, err := getLengthOfPath()
	if err != nil {
		return nil, err
	}

	for i := 0; i < resolvedLength; i++ {
		newOp := copyOperation(op)
		newPath := json.RawMessage(bytes.Replace(*newOp["path"], []byte(`/*/`), []byte(fmt.Sprintf("/%d/", i)), 1))
		newOp["path"] = &newPath
		resOps = append(resOps, newOp)
	}

	return resOps, nil
}
