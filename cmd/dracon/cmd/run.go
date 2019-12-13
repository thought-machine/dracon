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
	"os"

	"github.com/spf13/cobra"
	"github.com/thought-machine/dracon/pkg/kubernetes"
	"github.com/thought-machine/dracon/pkg/template"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run Dracon Pipeline",
	Long:  `Use run to execute a Dracon pipeline-run.`,
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

		// append PipelineResources
		pipelineResourceDocs, err := template.GeneratePipelineResourceDocs()
		files["draconPipelineResources"] = pipelineResourceDocs

		resDocs, err := template.PatchFileYAMLs(files, patches)
		if err != nil {
			return err
		}

		for _, doc := range resDocs["PipelineResource"] {
			err = kubernetes.Apply(string(doc))
			if err != nil {
				log.Fatalf("Failed to apply templates: %s\n", err)
				os.Exit(2)
			}
		}

		for _, doc := range resDocs["PipelineRun"] {
			err = kubernetes.Apply(string(doc))
			if err != nil {
				log.Fatalf("Failed to apply templates: %s\n", err)
				os.Exit(2)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVar(&pipelineOpts.PipelinePath, "pipeline", "", "Path to load the pipeline from")
	runCmd.Flags().StringVar(&pipelineOpts.ExtraPatchesPath, "extra-patches", "", "Path to load extra patches from")
	runCmd.MarkFlagRequired("pipeline")
}
