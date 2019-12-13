package putil

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
	v1 "github.com/thought-machine/dracon/pkg/genproto/v1"
)

// WriteEnrichedResults writes the given enriched results to the given output file
func WriteEnrichedResults(
	originalResults *v1.LaunchToolResponse,
	enrichedIssues []*v1.EnrichedIssue,
	outFile string,
) error {
	if err := os.MkdirAll(filepath.Dir(outFile), os.ModePerm); err != nil {
		return err
	}
	out := v1.EnrichedLaunchToolResponse{
		OriginalResults: originalResults,
		Issues:          enrichedIssues,
	}
	outBytes, err := proto.Marshal(&out)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(outFile, outBytes, 0644); err != nil {
		return err
	}

	log.Printf("wrote %d enriched issues to %s", len(enrichedIssues), outFile)
	return nil
}

// WriteResults writes the given issues to the the given output file as the given tool name
func WriteResults(
	toolName string,
	issues []*v1.Issue,
	outFile string,
) error {
	if err := os.MkdirAll(filepath.Dir(outFile), os.ModePerm); err != nil {
		return err
	}
	out := v1.LaunchToolResponse{
		ToolName: toolName,
		Issues:   issues,
	}

	outBytes, err := proto.Marshal(&out)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(outFile, outBytes, 0644); err != nil {
		return errors.Wrapf(err, "could not write to file %s", outFile)
	}

	log.Printf("wrote %d issues from to %s", len(issues), outFile)
	return nil
}
