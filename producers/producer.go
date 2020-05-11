// Package producers provides helper functions for writing Dracon compatible producers that parse tool outputs.
// Subdirectories in this package have more complete example usages of this package.
package producers

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	v1 "api/proto/v1"

	"github.com/thought-machine/dracon/pkg/putil"
)

var (
	InResults string
	OutFile   string
)

const (
	sourceDir = "/dracon/source"

	// EnvDraconStartTime Start Time of Dracon Scan in RFC3339
	EnvDraconStartTime = "DRACON_SCAN_TIME"
	// EnvDraconScanID the ID of the dracon scan
	EnvDraconScanID = "DRACON_SCAN_ID"
)

func init() {
	flag.StringVar(&InResults, "in", "", "")
	flag.StringVar(&OutFile, "out", "", "")
}

// ParseFlags will parse the input flags for the producer and perform simple validation
func ParseFlags() error {
	flag.Parse()
	if len(InResults) < 0 {
		return fmt.Errorf("in is undefined")
	}
	if len(OutFile) < 0 {
		return fmt.Errorf("out is undefined")
	}
	return nil
}

// ParseInFileJSON provides a generic method to parse a tool's JSON results into a given struct
func ParseInFileJSON(structure interface{}) error {
	inFile, err := os.Open(InResults)
	if err != nil {
		return err
	}
	dec := json.NewDecoder(inFile)
	for {
		if err := dec.Decode(structure); err == io.EOF {
			break
		} else if err != nil {
			return err
		}
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
