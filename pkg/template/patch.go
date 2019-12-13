package template

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/ghodss/yaml"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TypeYAMLDocs represents a list of resource docs per type in yaml format
type TypeYAMLDocs map[string][][]byte

type k8sDoc struct {
	metav1.TypeMeta `json:",inline"`
}

func getKindFromDoc(t []byte) (string, error) {
	var doc k8sDoc
	if err := json.Unmarshal(t, &doc); err != nil {
		return "", err
	}

	return doc.GroupVersionKind().Kind, nil
}

// PatchFileYAMLs applies the given patches to files
func PatchFileYAMLs(
	files PipelineYAMLDocs,
	patches PatchKindYAMLDocs,
) (TypeYAMLDocs, error) {
	resDocs := TypeYAMLDocs{}
	for path, f := range files {
		log.Printf("processing: %s", path)
		for _, t := range f {
			patchKind, err := getKindFromDoc(t)
			if err != nil {
				return nil, fmt.Errorf("could not get kind from doc %s: %w", t, err)
			}
			if _, ok := resDocs[patchKind]; !ok {
				resDocs[patchKind] = [][]byte{}
			}
			buf := &bytes.Buffer{}
			yamlBytes, err := yaml.JSONToYAML(t)
			if err != nil {
				return nil, fmt.Errorf("could not translate from JSON to YAML %s: %w", t, err)
			}
			if foundPatches, ok := patches[patchKind]; ok {
				log.Printf("applying patch kind: %s", patchKind)
				modifiedT := t
				for _, patch := range foundPatches {
					newPatch := jsonpatch.Patch{}
					for _, op := range patch {
						globPatch, err := patchArrayGlob(op, t)
						if err != nil {
							return nil, err
						}
						newPatch = append(newPatch, globPatch...)
					}
					modifiedT, err = newPatch.Apply(modifiedT)
					if err != nil {
						return nil, err
					}
					yamlModified, err := yaml.JSONToYAML(modifiedT)
					if err != nil {
						return nil, err
					}
					yamlBytes = yamlModified
				}
			}
			buf.WriteString(fmt.Sprintf("---\n%s\n", yamlBytes))
			resDocs[patchKind] = append(resDocs[patchKind], buf.Bytes())
		}
	}
	return resDocs, nil
}

func copyOperation(op jsonpatch.Operation) jsonpatch.Operation {
	newOp := jsonpatch.Operation{}
	newOp["op"] = op["op"]
	if _, ok := op["path"]; ok {
		newOp["path"] = op["path"]
	}
	if _, ok := op["value"]; ok {
		newOp["value"] = op["value"]
	}
	if _, ok := op["from"]; ok {
		newOp["from"] = op["from"]
	}
	return newOp
}
