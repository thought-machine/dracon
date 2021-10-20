// Package producers provides helper functions for writing Dracon compatible producers that parse tool outputs.
// Subdirectories in this package have more complete example usages of this package.
package producers

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	v1 "github.com/thought-machine/dracon/api/proto/v1"

	"github.com/thought-machine/dracon/pkg/putil"
)

var (
	// InResults represents incoming tool output
	InResults string
	// OutFile points to the protobuf file where dracon results will be written
	OutFile string
	// Append flag will append to the outfile instead of overwriting, useful when there's multiple inresults
	Append bool
)

const (
	sourceDir = "/dracon/source"

	// EnvDraconStartTime Start Time of Dracon Scan in RFC3339
	EnvDraconStartTime = "DRACON_SCAN_TIME"
	// EnvDraconScanID the ID of the dracon scan
	EnvDraconScanID = "DRACON_SCAN_ID"
)

// ParseFlags will parse the input flags for the producer and perform simple validation
func ParseFlags() error {
	flag.StringVar(&InResults, "in", "", "")
	flag.StringVar(&OutFile, "out", "", "")
	flag.BoolVar(&Append, "append", false, "Append to output file instead of overwriting it")

	flag.Parse()
	if len(InResults) < 0 {
		return fmt.Errorf("in is undefined")
	}
	if len(OutFile) < 0 {
		return fmt.Errorf("out is undefined")
	}
	return nil
}

// ReadInFile returns the contents of the file given by InResults.
func ReadInFile() ([]byte, error) {
	file, err := os.Open(InResults)
	if err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(file)
	return buffer.Bytes(), nil
}

// ParseJSON provides a generic method to parse JSON input (e.g. the results
// provided by a tool) into a given struct.
func ParseJSON(in []byte, structure interface{}) error {
	if err := json.Unmarshal(in, &structure); err != nil {
		return err
	}
	return nil
}

// WriteDraconOut provides a generic method to write the resulting protobuf to the output file
func WriteDraconOut(
	toolName string,
	issues []*v1.Issue,
) error {
	source := getSource()
	cleanIssues := []*v1.Issue{}
	for _, iss := range issues {
		iss.Description = strings.Replace(iss.Description, sourceDir, ".", -1)
		iss.Title = strings.Replace(iss.Title, sourceDir, ".", -1)
		iss.Target = strings.Replace(iss.Target, sourceDir, ".", -1)
		iss.Source = source
		cleanIssues = append(cleanIssues, iss)
		log.Printf("found issue: %+v\n", iss)
	}
	scanStartTime := os.Getenv(EnvDraconStartTime)
	if scanStartTime == "" {
		scanStartTime = time.Now().UTC().Format(time.RFC3339)
	}
	scanUUUID := os.Getenv(EnvDraconScanID)

	stat, err := os.Stat(OutFile)
	if Append && err == nil && stat.Size() > 0 {
		return putil.AppendResults(cleanIssues, OutFile)
	}
	return putil.WriteResults(toolName, cleanIssues, OutFile, scanUUUID, scanStartTime)
}

func getSource() string {
	sourceMetaPath := filepath.Join(sourceDir, ".source.dracon")
	_, err := os.Stat(sourceMetaPath)
	if os.IsNotExist(err) {
		return "unknown"
	}

	dat, err := ioutil.ReadFile(sourceMetaPath)
	if err != nil {
		log.Println(err)
	}
	return strings.TrimSpace(string(dat))
}
