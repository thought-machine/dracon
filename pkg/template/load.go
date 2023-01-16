package template

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"io/fs"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/ghodss/yaml"
	goyaml "gopkg.in/yaml.v3"
)

// yamlPatches holds patch yaml files as static assets
//go:embed *.yaml
var yamlPatches embed.FS

// Errors returned from this package
var (
	ErrNonYAMLFileEncountered = errors.New("non-yaml file found in directory")
)

// PipelineYAMLDocs stores all of the yaml docs found in a file in the format map[path][]doc
type PipelineYAMLDocs map[string]ResourceDocs

// LoadPipelineYAMLFiles returns all of the PipelineYAMLDocs in a directory
func LoadPipelineYAMLFiles(sourcePath string) (PipelineYAMLDocs, error) {
	targets := PipelineYAMLDocs{}
	err := filepath.Walk(sourcePath, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() && (strings.HasSuffix(f.Name(), ".yml") || strings.HasSuffix(f.Name(), ".yaml")) {
			docs, err := loadYAMLFile(path)
			if err != nil {
				return err
			}
			targets[path] = docs
		}
		return nil
	})
	return targets, err
}

func loadYAMLFile(path string) (ResourceDocs, error) {
	targetYAML, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read file at path %s: %w", path, err)
	}
	resFileYamlDocs := ResourceDocs{}
	yamlByteParts, err := splitYAML(targetYAML)
	if err != nil {
		return nil, fmt.Errorf("could not split YAML at path %s: %w", path, err)
	}
	yamlParts := ResourceDocs{}
	for _, bytePart := range yamlByteParts {
		yamlParts = append(yamlParts, ResourceDoc(bytePart))
	}
	yamlDocs := func(yamlParts ResourceDocs) ResourceDocs {
		yamlDocs := ResourceDocs{}
		for _, d := range yamlParts {
			if strings.TrimSpace(string(d)) != "" {
				yamlDocs = append(yamlDocs, d)
			}
		}
		return yamlDocs
	}(yamlParts)
	log.Printf("found %d YAML docs in %s", len(yamlDocs), path)
	for _, yDoc := range yamlDocs {
		yDocParsed, err := yaml.YAMLToJSON(yDoc)
		if err != nil {
			return nil, fmt.Errorf("could not read YAML doc in path %s: %w", path, err)
		}
		resFileYamlDocs = append(resFileYamlDocs, yDocParsed)
	}

	return resFileYamlDocs, nil
}

func splitYAML(targetYAML []byte) ([][]byte, error) {
	dec := goyaml.NewDecoder(bytes.NewReader(targetYAML))

	var res [][]byte
	for {
		var value interface{}
		err := dec.Decode(&value)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		valueBytes, err := goyaml.Marshal(value)
		if err != nil {
			return nil, err
		}
		res = append(res, valueBytes)
	}
	return res, nil
}

// PatchKindYAMLDocs stores all of the jsonpatch yaml docs found by type
type PatchKindYAMLDocs map[string][]jsonpatch.Patch

func loadStatikPatches() (PatchKindYAMLDocs, error) {
	patches := PatchKindYAMLDocs{}
	err := fs.WalkDir(yamlPatches, ".", func(path string, f fs.DirEntry, err error) error {
		if !f.IsDir() {
			patchKind := getPatchKindFromPath(path)
			contents, err := fs.ReadFile(yamlPatches, path)
			if err != nil {
				return fmt.Errorf("could not read file: %w", err)
			}
			patch, err := loadPatchFromYAML(contents)
			if err != nil {
				return fmt.Errorf("could not load patch from YAML: %w", err)
			}
			if _, ok := patches[patchKind]; !ok {
				patches[patchKind] = []jsonpatch.Patch{}
			}
			patches[patchKind] = append(patches[patchKind], patch)
		}
		return nil
	})

	return patches, err
}

// LoadPatchYAMLFiles returns the yaml docs by kind from a given directory
func LoadPatchYAMLFiles(sourcePath string) (PatchKindYAMLDocs, error) {
	patches, err := loadStatikPatches()
	if err != nil {
		return nil, err
	}
	if sourcePath != "" {
		err = filepath.Walk(sourcePath, func(path string, f os.FileInfo, err error) error {
			if !f.IsDir() && (strings.HasSuffix(f.Name(), ".yml") || strings.HasSuffix(f.Name(), ".yaml")) {
				patchKind := getPatchKindFromPath(path)
				patchYAML, err := ioutil.ReadFile(path)
				if err != nil {
					return fmt.Errorf("could not read file: %w", err)
				}
				patch, err := loadPatchFromYAML(patchYAML)
				if err != nil {
					return fmt.Errorf("could not load patch from YAML: %w", err)
				}
				if _, ok := patches[patchKind]; !ok {
					patches[patchKind] = []jsonpatch.Patch{}
				}
				patches[patchKind] = append(patches[patchKind], patch)
			}
			return nil
		})
	}
	return patches, err
}

// getPatchKindFromPath returns the type of yaml file based on filename
func getPatchKindFromPath(path string) string {
	base := filepath.Base(path)
	parts := strings.Split(base, `.`)
	return parts[len(parts)-2]
}

func loadPatchFromYAML(contents []byte) (jsonpatch.Patch, error) {
	templatedPatchYAML, err := execTemplate(contents)
	if err != nil {
		return nil, err
	}
	patchJSON, err := yaml.YAMLToJSON(templatedPatchYAML)
	if err != nil {
		return nil, err
	}
	return jsonpatch.DecodePatch(patchJSON)
}
