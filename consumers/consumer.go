// Package consumers provides helper functions for working with Dracon compatible outputs as a Consumer.
// Subdirectories in this package have more complete example usages of this package.
package consumers

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/golang/protobuf/ptypes"
	v1 "github.com/thought-machine/dracon/pkg/genproto/v1"
	"github.com/thought-machine/dracon/pkg/putil"
)

const (
	// EnvDraconStartTime Start Time of Dracon Scan in RFC3339
	EnvDraconStartTime = "DRACON_SCAN_TIME"
	// EnvDraconScanID the ID of the dracon scan
	EnvDraconScanID = "DRACON_SCAN_ID"
)

var (
	inResults string
	// Raw represents if the non-enriched results should be used
	Raw bool
)

func init() {
	flag.StringVar(&inResults, "in", "", "the directory where dracon producer/enricher outputs are")
	flag.BoolVar(&Raw, "raw", false, "if the non-enriched results should be used")
}

// ParseFlags will parse the input flags for the producer and perform simple validation
func ParseFlags() error {
	flag.Parse()
	if len(inResults) < 1 {
		return fmt.Errorf("in is undefined")
	}
	return nil
}

// LoadToolResponse loads raw results from producers
func LoadToolResponse() ([]*v1.LaunchToolResponse, error) {
	res, err := putil.LoadToolResponse(inResults)
	if err != nil {
		return nil, err
	}
	scanInfo, err := getScanInfo()
	if err != nil {
		return nil, err
	}
	for _, r := range res {
		r.ScanInfo = scanInfo
	}
	return res, nil
}

// LoadEnrichedToolResponse loads enriched results from the enricher
func LoadEnrichedToolResponse() ([]*v1.EnrichedLaunchToolResponse, error) {
	res, err := putil.LoadEnrichedToolResponse(inResults)
	if err != nil {
		return nil, err
	}
	scanInfo, err := getScanInfo()
	if err != nil {
		return nil, err
	}
	for _, r := range res {
		r.OriginalResults.ScanInfo = scanInfo
	}
	return res, nil
}

func getScanInfo() (*v1.ScanInfo, error) {
	t, err := time.Parse(time.RFC3339, os.Getenv(EnvDraconStartTime))
	if err != nil {
		return nil, err
	}
	tp, err := ptypes.TimestampProto(t)
	if err != nil {
		return nil, err
	}
	return &v1.ScanInfo{
		ScanUuid:      os.Getenv(EnvDraconScanID),
		ScanStartTime: tp,
	}, nil
}
