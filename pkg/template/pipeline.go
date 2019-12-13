package template

// non-runtime
import (
	"encoding/json"
	"strings"
)

// PipelineParam represents a Parameter in a Pipeline
type PipelineParam struct {
	Name        string
	Description string
	Type        string

	Value string
}

// PipelineTask represents a Task in a Pipeline
type PipelineTask struct {
	Index int
	Name  string
}

func preparePipelineVars(doc []byte) error {
	var err error
	TemplateVars.PipelineTaskProducers, err = getPipelineTasksByType("producer", doc)
	if err != nil {
		return err
	}
	TemplateVars.PipelineTaskConsumers, err = getPipelineTasksByType("consumer", doc)
	if err != nil {
		return err
	}
	TemplateVars.PipelineTaskEnrichers, err = getPipelineTasksByType("enricher", doc)
	if err != nil {
		return err
	}
	return nil
}

func getPipelineTasksByType(taskType string, targetJSON []byte) ([]PipelineTask, error) {
	var p pipeline
	err := json.Unmarshal(targetJSON, &p)
	if err != nil {
		return nil, err
	}

	pipelineTasks := []PipelineTask{}
	for i, t := range p.Spec.Tasks {
		tTaskType := t.Name
		taskName := t.Name
		nameParts := strings.Split(t.Name, "-")
		if len(nameParts) > 1 {
			tTaskType = nameParts[len(nameParts)-1]
			taskName = strings.Join(nameParts[0:len(nameParts)-1], "-")
		}
		if tTaskType == taskType {
			pipelineTasks = append(pipelineTasks, PipelineTask{Index: i, Name: taskName})
		}
	}

	return pipelineTasks, nil
}

type pipeline struct {
	Spec struct {
		Tasks []struct {
			Name string `json:"name"`
		} `json:"tasks"`
	} `json:"spec"`
}
