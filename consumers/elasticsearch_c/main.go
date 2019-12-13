package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	//  TODO(hjenkins): Support multiple versions of ES
	// elasticsearchv5 "github.com/elastic/go-elasticsearch/v5"
	elasticsearchv6 "github.com/elastic/go-elasticsearch"
	// elasticsearchv7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/golang/protobuf/ptypes"
	"github.com/thought-machine/dracon/consumers"
	v1 "github.com/thought-machine/dracon/pkg/genproto/v1"
)

var (
	esURL   string
	esIndex string
)

func init() {
	flag.StringVar(&esIndex, "es-index", "", "the index in elasticsearch to push results to")
}

func parseFlags() error {
	if err := consumers.ParseFlags(); err != nil {
		return err
	}
	if len(esIndex) < 1 {
		return fmt.Errorf("es-index is undefined")
	}
	return nil
}

func main() {
	if err := consumers.ParseFlags(); err != nil {
		log.Fatal(err)
	}

	if err := getESClient(); err != nil {
		log.Fatal(err)
	}

	if consumers.Raw {
		responses, err := consumers.LoadToolResponse()
		if err != nil {
			log.Fatal(err)
		}
		for _, res := range responses {
			scanStartTime, _ := ptypes.Timestamp(res.GetScanInfo().GetScanStartTime())
			for _, iss := range res.GetIssues() {
				b, err := getRawIssue(scanStartTime, res, iss)
				if err != nil {
					log.Fatal(err)
				}
				esPush(b)
			}
		}
	} else {
		responses, err := consumers.LoadEnrichedToolResponse()
		if err != nil {
			log.Fatal(err)
		}
		for _, res := range responses {
			scanStartTime, _ := ptypes.Timestamp(res.GetOriginalResults().GetScanInfo().GetScanStartTime())
			for _, iss := range res.GetIssues() {
				b, err := getEnrichedIssue(scanStartTime, res, iss)
				if err != nil {
					log.Fatal(err)
				}
				esPush(b)
			}
		}
	}
}

func getRawIssue(scanStartTime time.Time, res *v1.LaunchToolResponse, iss *v1.Issue) ([]byte, error) {
	jBytes, err := json.Marshal(&esDocument{
		ScanStartTime: scanStartTime,
		ScanID:        res.GetScanInfo().GetScanUuid(),
		ToolName:      res.GetToolName(),
		Source:        iss.GetSource(),
		Title:         iss.GetTitle(),
		Target:        iss.GetTarget(),
		Type:          iss.GetType(),
		Severity:      iss.GetSeverity(),
		CVSS:          iss.GetCvss(),
		Confidence:    iss.GetConfidence(),
		Description:   iss.GetDescription(),
		FirstFound:    scanStartTime,
		FalsePositive: false,
	})
	if err != nil {
		return []byte{}, err
	}
	return jBytes, nil
}

func getEnrichedIssue(scanStartTime time.Time, res *v1.EnrichedLaunchToolResponse, iss *v1.EnrichedIssue) ([]byte, error) {
	firstSeenTime, _ := ptypes.Timestamp(iss.GetFirstSeen())
	jBytes, err := json.Marshal(&esDocument{
		ScanStartTime: scanStartTime,
		ScanID:        res.GetOriginalResults().GetScanInfo().GetScanUuid(),
		ToolName:      res.GetOriginalResults().GetToolName(),
		Source:        iss.GetRawIssue().GetSource(),
		Title:         iss.GetRawIssue().GetTitle(),
		Target:        iss.GetRawIssue().GetTarget(),
		Type:          iss.GetRawIssue().GetType(),
		Severity:      iss.GetRawIssue().GetSeverity(),
		CVSS:          iss.GetRawIssue().GetCvss(),
		Confidence:    iss.GetRawIssue().GetConfidence(),
		Description:   iss.GetRawIssue().GetDescription(),
		FirstFound:    firstSeenTime,
		FalsePositive: iss.GetFalsePositive(),
	})
	if err != nil {
		return []byte{}, err
	}
	return jBytes, nil
}

type esDocument struct {
	ScanStartTime time.Time     `json:"scan_start_time"`
	ScanID        string        `json:"scan_id"`
	ToolName      string        `json:"tool_name"`
	Source        string        `json:"source"`
	Target        string        `json:"target"`
	Type          string        `json:"type"`
	Title         string        `json:"title"`
	Severity      v1.Severity   `json:"severity"`
	CVSS          float64       `json:"cvss"`
	Confidence    v1.Confidence `json:"confidence"`
	Description   string        `json:"description"`
	FirstFound    time.Time     `json:"first_found"`
	FalsePositive bool          `json:"false_positive"`
}

var esClient interface{}

func getESClient() error {
	es, err := elasticsearchv6.NewDefaultClient()
	if err != nil {
		return err
	}

	type esInfo struct {
		Version struct {
			Number string `json:"number"`
		} `json:"version"`
	}

	res, err := es.Info()
	if err != nil {
		return err
	}
	var info esInfo
	if err := json.NewDecoder(res.Body).Decode(&info); err != nil {
		return err
	}
	switch info.Version.Number[0] {
	// case '5':
	// 	esClient, err = elasticsearchv5.NewDefaultClient()
	case '6':
		esClient, err = elasticsearchv6.NewDefaultClient()
	// case '7':
	// 	esClient, err = elasticsearchv7.NewDefaultClient()
	default:
		err = fmt.Errorf("unsupported version %s", info.Version.Number)
	}
	return err
}

func esPush(b []byte) error {
	var err error
	var res interface{}
	switch x := esClient.(type) {
	// case *elasticsearchv5.Client:
	// 	res, err = x.Index(esIndex, bytes.NewBuffer(b), x.Index.WithDocumentType("doc"))
	case *elasticsearchv6.Client:
		res, err = x.Index(esIndex, bytes.NewBuffer(b), x.Index.WithDocumentType("doc"))
	// case *elasticsearchv7.Client:
	// 	res, err = x.Index(esIndex, bytes.NewBuffer(b))
	default:
		err = fmt.Errorf("unsupported client %T", esClient)
	}
	log.Printf("%+v", res)
	return err
}
