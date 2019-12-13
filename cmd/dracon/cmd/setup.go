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
package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/thought-machine/dracon/pkg/kubernetes"
	"github.com/thought-machine/dracon/pkg/template"
)

var pipelineOpts struct {
	PipelinePath     string
	ExtraPatchesPath string
}

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup a new Dracon Pipeline",
	Long:  `Use setup to help with setting up a new Dracon pipeline.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// load pipeline files
		files, err := template.LoadPipelineYAMLFiles(pipelineOpts.PipelinePath)
		if err != nil {
			return err
		}
		// prepare template vars from target files and patch files
		if err := template.PrepareVars(files); err != nil {
			return err
		}
		// apply template to target files
		files, err = template.ExecuteFiles(files)
		if err != nil {
			return err
		}

		// load all patch files
		patches, err := template.LoadPatchYAMLFiles(pipelineOpts.ExtraPatchesPath)
		if err != nil {
			return err
		}

		resDocs, err := template.PatchFileYAMLs(files, patches)
		if err != nil {
			return err
		}

		for k, docs := range resDocs {
			if k != "PipelineRun" && k != "PipelineResource" {
				for _, doc := range docs {
					err = kubernetes.Apply(string(doc))
					if err != nil {
						log.Fatalf("Failed to apply templates :%s\n", err)
					}
				}
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)

	setupCmd.Flags().StringVar(&pipelineOpts.PipelinePath, "pipeline", "", "Path to load the pipeline from")
	setupCmd.Flags().StringVar(&pipelineOpts.ExtraPatchesPath, "extra-patches", "", "Path to load extra patches from")
	setupCmd.MarkFlagRequired("pipeline")
}
