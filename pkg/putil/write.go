package putil

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	v1 "github.com/thought-machine/dracon/api/proto/v1"
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
	scanUUID string,
	scanStartTime string,
) error {
	if err := os.MkdirAll(filepath.Dir(outFile), os.ModePerm); err != nil {
		return err
	}
	timeVal, err := time.Parse(time.RFC3339, scanStartTime)
	if err != nil {
		return err
	}
	timestamp, err := ptypes.TimestampProto(timeVal)
	if err != nil {
		return err
	}
	scanInfo := v1.ScanInfo{
		ScanUuid:      scanUUID,
		ScanStartTime: timestamp,
	}
	out := v1.LaunchToolResponse{
		ScanInfo: &scanInfo,
		ToolName: toolName,
		Issues:   issues,
	}

	outBytes, err := proto.Marshal(&out)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(outFile, outBytes, 0644); err != nil {
		return fmt.Errorf("could not write to file '%s': %w", outFile, err)
	}

	log.Printf("wrote %d issues from to %s", len(issues), outFile)
	return nil
}
