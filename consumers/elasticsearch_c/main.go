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
	// elasticsearchv6 "github.com/elastic/go-elasticsearch/v6"
	"api/proto/v1"

	elasticsearchv7 "github.com/elastic/go-elasticsearch"
	"github.com/golang/protobuf/ptypes"
	"github.com/thought-machine/dracon/consumers"
)

var (
	esURL         string
	esIndex       string
	basicAuthUser string
	basicAuthPass string
)

func init() {
	flag.StringVar(&esIndex, "es-index", "", "the index in elasticsearch to push results to")
	flag.StringVar(&basicAuthUser, "basic-auth-user", "", "[OPTIONAL] the basic auth username")
	flag.StringVar(&basicAuthPass, "basic-auth-pass", "", "[OPTIONAL] the basic auth password")

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
		log.Fatal("could not contact remote Elasticsearch: ", err)
	}

	if consumers.Raw {
		log.Print("Parsing Raw results")
		responses, err := consumers.LoadToolResponse()
		if err != nil {
			log.Fatal("could not load raw results, file malformed: ", err)
		}
		for _, res := range responses {
			scanStartTime, _ := ptypes.Timestamp(res.GetScanInfo().GetScanStartTime())
			for _, iss := range res.GetIssues() {
				fmt.Printf("Pushing %d, issues to es \n", len(responses))
				b, err := getRawIssue(scanStartTime, res, iss)
				if err != nil {
					log.Fatal("Could not parse raw issue", err)
				}
				esPush(b)
			}
		}
	} else {
		log.Print("Parsing Enriched results")
		responses, err := consumers.LoadEnrichedToolResponse()
		if err != nil {
			log.Fatal("could not load enriched results, file malformed: ", err)
		}
		for _, res := range responses {
			scanStartTime, _ := ptypes.Timestamp(res.GetOriginalResults().GetScanInfo().GetScanStartTime())
			for _, iss := range res.GetIssues() {
				b, err := getEnrichedIssue(scanStartTime, res, iss)
				if err != nil {
					log.Fatal("Could not parse enriched issue", err)
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
func severtiyToText(severity v1.Severity) string {
	switch severity {
	case v1.Severity_SEVERITY_INFO:
		return "Info"
	case v1.Severity_SEVERITY_LOW:
		return "Low"
	case v1.Severity_SEVERITY_MEDIUM:
		return "Medium"
	case v1.Severity_SEVERITY_HIGH:
		return "High"
	case v1.Severity_SEVERITY_CRITICAL:
		return "Critical"
	default:
		return "N/A"
	}
}
func confidenceToText(confidence v1.Confidence) string {
	switch confidence {
	case v1.Confidence_CONFIDENCE_INFO:
		return "Info"
	case v1.Confidence_CONFIDENCE_LOW:
		return "Low"
	case v1.Confidence_CONFIDENCE_MEDIUM:
		return "Medium"
	case v1.Confidence_CONFIDENCE_HIGH:
		return "High"
	case v1.Confidence_CONFIDENCE_CRITICAL:
		return "Critical"
	default:
		return "N/A"
	}

}
func getEnrichedIssue(scanStartTime time.Time, res *v1.EnrichedLaunchToolResponse, iss *v1.EnrichedIssue) ([]byte, error) {
	firstSeenTime, _ := ptypes.Timestamp(iss.GetFirstSeen())
	jBytes, err := json.Marshal(&esDocument{
		ScanStartTime:  scanStartTime,
		ScanID:         res.GetOriginalResults().GetScanInfo().GetScanUuid(),
		ToolName:       res.GetOriginalResults().GetToolName(),
		Source:         iss.GetRawIssue().GetSource(),
		Title:          iss.GetRawIssue().GetTitle(),
		Target:         iss.GetRawIssue().GetTarget(),
		Type:           iss.GetRawIssue().GetType(),
		Severity:       iss.GetRawIssue().GetSeverity(),
		CVSS:           iss.GetRawIssue().GetCvss(),
		Confidence:     iss.GetRawIssue().GetConfidence(),
		Description:    iss.GetRawIssue().GetDescription(),
		FirstFound:     firstSeenTime,
		FalsePositive:  iss.GetFalsePositive(),
		SeverityText:   severtiyToText(iss.GetRawIssue().GetSeverity()),
		ConfidenceText: confidenceToText(iss.GetRawIssue().GetConfidence()),
	})
	if err != nil {
		return []byte{}, err
	}
	return jBytes, nil
}

type esDocument struct {
	ScanStartTime  time.Time   `json:"scan_start_time"`
	ScanID         string      `json:"scan_id"`
	ToolName       string      `json:"tool_name"`
	Source         string      `json:"source"`
	Target         string      `json:"target"`
	Type           string      `json:"type"`
	Title          string      `json:"title"`
	Severity       v1.Severity `json:"severity"`
	SeverityText   string      `json:"severity_text"`
	CVSS           float64       `json:"cvss"`
	Confidence     v1.Confidence `json:"confidence"`
	ConfidenceText string 	 `json:"confidence_text"`
	Description    string    `json:"description"`
	FirstFound     time.Time `json:"first_found"`
	FalsePositive  bool      `json:"false_positive"`
}

var esClient interface{}

func getESClient() error {
	var es *elasticsearchv7.Client
	var err error = nil
	if basicAuthUser != "" && basicAuthPass != "" {
		es, err = elasticsearchv7.NewClient(elasticsearchv7.Config{
			Username: basicAuthUser,
			Password: basicAuthPass,
		})
	} else {
		es, err = elasticsearchv7.NewDefaultClient()
	}
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
	// case '6':
	// 	esClient, err = elasticsearchv6.NewDefaultClient()
	case '7':
		if basicAuthUser != "" && basicAuthPass != "" {
			esClient, err = elasticsearchv7.NewClient(elasticsearchv7.Config{
				Username: basicAuthUser,
				Password: basicAuthPass,
			})
		} else {
			esClient, err = elasticsearchv7.NewDefaultClient()
		}

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
	// case *elasticsearchv6.Client:
	// 	res, err = x.Index(esIndex, bytes.NewBuffer(b), x.Index.WithDocumentType("doc"))
	case *elasticsearchv7.Client:
		res, err = x.Index(esIndex, bytes.NewBuffer(b))
	default:
		err = fmt.Errorf("unsupported client %T", esClient)
	}
	log.Printf("%+v", res)
	return err
}
