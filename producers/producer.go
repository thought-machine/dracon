// Package producers provides helper functions for writing Dracon compatible producers that parse tool outputs.
// Subdirectories in this package have more complete example usages of this package.
package producers

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	v1 "github.com/thought-machine/dracon/pkg/genproto/v1"
	"github.com/thought-machine/dracon/pkg/putil"
)

var (
	inResults string
	outFile   string
)

const sourceDir = "/dracon/source"

func init() {
	flag.StringVar(&inResults, "in", "", "")
	flag.StringVar(&outFile, "out", "", "")
}

// ParseFlags will parse the input flags for the producer and perform simple validation
func ParseFlags() error {
	flag.Parse()
	if len(inResults) < 0 {
		return fmt.Errorf("in is undefined")
	}
	if len(outFile) < 0 {
		return fmt.Errorf("out is undefined")
	}
	return nil
}

// ParseInFileJSON provides a generic method to parse a tool's JSON results into a given struct
func ParseInFileJSON(structure interface{}) error {
	inBytes, err := ioutil.ReadFile(inResults)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(inBytes, structure); err != nil {
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

	return putil.WriteResults(toolName, cleanIssues, outFile)
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
