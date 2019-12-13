package putil

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gogo/protobuf/proto"
	v1 "github.com/thought-machine/dracon/pkg/genproto/v1"
)

// LoadToolResponse loads raw results
func LoadToolResponse(inPath string) ([]*v1.LaunchToolResponse, error) {
	responses := []*v1.LaunchToolResponse{}
	if err := filepath.Walk(inPath, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() && (strings.HasSuffix(f.Name(), ".pb")) {
			pbBytes, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			res := v1.LaunchToolResponse{}
			if err := proto.Unmarshal(pbBytes, &res); err != nil {
				log.Printf("skipping %s as unable to unmarshal", path)
			} else {
				responses = append(responses, &res)
			}
		}
		return nil
	}); err != nil {
		return responses, err
	}
	return responses, nil
}

// LoadEnrichedToolResponse loads enriched results from the enricher
func LoadEnrichedToolResponse(inPath string) ([]*v1.EnrichedLaunchToolResponse, error) {
	responses := []*v1.EnrichedLaunchToolResponse{}
	if err := filepath.Walk(inPath, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() && (strings.HasSuffix(f.Name(), ".pb")) {
			pbBytes, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			res := v1.EnrichedLaunchToolResponse{}
			if err := proto.Unmarshal(pbBytes, &res); err != nil {
				log.Printf("skipping %s as unable to unmarshal", path)
			} else {
				responses = append(responses, &res)
			}
		}
		return nil
	}); err != nil {
		return responses, err
	}
	return responses, nil
}
