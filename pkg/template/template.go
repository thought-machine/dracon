/*
Copyright © 2019 Thought Machine

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package template

import (
	"bytes"
	"fmt"
	"html/template"
	"time"

	"github.com/speps/go-hashids"
)

// ResourceDoc represents a K8s resource document
type ResourceDoc []byte

// ResourceDocs represents a set of K8s resource documents
type ResourceDocs []ResourceDoc

// TemplateVars represents the vars that are available in all templates
var TemplateVars = newTemplateVars()

// PrepareVars adds contextual vars to the templater
func PrepareVars(files PipelineYAMLDocs) error {
	for _, f := range files {
		for _, t := range f {
			patchKind, err := getKindFromDoc(t)
			if err != nil {
				return err
			}
			switch patchKind {
			case "Pipeline":
				err = preparePipelineVars(t)
			}
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// ExecuteFiles executes the templates
func ExecuteFiles(files PipelineYAMLDocs) (PipelineYAMLDocs, error) {
	templatedFiles := map[string]ResourceDocs{}
	for path, file := range files {
		templatedFiles[path] = ResourceDocs{}
		for _, target := range file {
			templatedTarget, err := execTemplate(target)
			if err != nil {
				return nil, err
			}
			templatedFiles[path] = append(templatedFiles[path], templatedTarget)
		}
	}
	return templatedFiles, nil
}

func execTemplate(targetJSON []byte) ([]byte, error) {
	t := template.Must(template.New("target").Parse(string(targetJSON)))
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, TemplateVars); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type templateVars struct {
	RunID string

	ProducerSourcePath  string
	ProducerToolOutPath string
	ProducerOutPath     string
	EnricherOutPath     string
	ConsumerSourcePath  string

	PipelineParams        []PipelineParam
	PipelineTaskEnrichers []PipelineTask
	PipelineTaskProducers []PipelineTask
	PipelineTaskConsumers []PipelineTask
}

func newTemplateVars() *templateVars {
	id := getID()
	return &templateVars{
		RunID:               id,
		ProducerSourcePath:  `/dracon/source`,
		ProducerToolOutPath: `/dracon/results`,
		ProducerOutPath:     `/workspace/output/producer/results.pb`,
		EnricherOutPath:     `/workspace/output/enricher`,
		ConsumerSourcePath:  `/workspace/`,
		PipelineParams: []PipelineParam{
			PipelineParam{
				"DRACON_SCAN_ID",
				"Dracon: Unique Scan ID",
				"string",
				fmt.Sprintf("dracon-%s", id),
			},
			PipelineParam{
				"DRACON_SCAN_TIME",
				"Dracon: Scan start time",
				"string",
				time.Now().UTC().Format(time.RFC3339),
			},
		},
	}
}

func getID() string {
	hd := hashids.NewData()
	hd.Alphabet = "abcdefghijklmnopqrstuvwxyz1234567890"
	hd.Salt = "dracon"
	hd.MinLength = 4
	h, _ := hashids.NewWithData(hd)
	e, _ := h.EncodeInt64([]int64{time.Now().UnixNano()})
	return e
}
