package main

import (
	v1 "github.com/thought-machine/dracon/api/proto/v1"
	types "github.com/thought-machine/dracon/producers/semgrep/types/semgrep-issue"

	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const exampleOutput = `
{
	"results": [
	  	{
			"check_id": "rules.go.xss.Go using template.HTML", 
			"path": "/src/go/xss/template-html.go",
			"start": {"line": 10, "col": 11},
			"end": {"line": 10, "col": 32},
			"extra": {
				"message": "Use of this type presents a security risk: the encapsulated content should come from a trusted source, \nas it will be included verbatim in the template output.\nhttps://blogtitle.github.io/go-safe-html/\n", 
				"metavars": {},
				"metadata": {}, 
				"severity": "WARNING", 
				"lines": "\t\t\treturn template.HTML(revStr)"
			}
		},
		{
			"check_id": "rules.python.grpc.GRPC Insecure Port",
			"path": "/src/python/grpc/grpc_insecure_port.py",
			"start": {"line": 19, "col": 5},
			"end": {"line": 19, "col": 68},
			"extra": {
				"message": "The gRPC server listening port is configured insecurely, this offers no encryption and authentication.\nPlease review and ensure that this is appropriate for the communication.  \n", 
				"metavars": {
					"$VAR": {
						"start": {"line": 19, "col": 5, "offset": 389}, 
						"end": {"line": 19, "col": 20, "offset": 404},
						"abstract_content": "insecure_server",
						"unique_id": {
							"type": "id", "value": "insecure_server",
							"kind": "Local", "sid": 8
						}
					}
				},
				"metadata": {},
				"severity": "WARNING",
				"lines": "    insecure_server.add_insecure_port('[::]:{}'.format(flags.port))"
			}
 		}
	]
}
`

func TestParseIssues(t *testing.T) {
	semgrepResults := types.SemgrepResults{}
	err := json.Unmarshal([]byte(exampleOutput), &semgrepResults)

	assert.Nil(t, err)
	issues := parseIssues(semgrepResults)

	expectedIssue := &v1.Issue{
		Target:      "/src/go/xss/template-html.go:10-10",
		Type:        "Use of this type presents a security risk: the encapsulated content should come from a trusted source, \nas it will be included verbatim in the template output.\nhttps://blogtitle.github.io/go-safe-html/\n",
		Title:       "rules.go.xss.Go using template.HTML",
		Severity:    v1.Severity_SEVERITY_MEDIUM,
		Cvss:        0.0,
		Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
		Description: "\t\t\treturn template.HTML(revStr)",
	}

	assert.Equal(t, expectedIssue, issues[0])

	expectedIssue2 := &v1.Issue{
		Target:      "/src/python/grpc/grpc_insecure_port.py:19-19",
		Type:        "The gRPC server listening port is configured insecurely, this offers no encryption and authentication.\nPlease review and ensure that this is appropriate for the communication.  \n",
		Title:       "rules.python.grpc.GRPC Insecure Port",
		Severity:    v1.Severity_SEVERITY_MEDIUM,
		Cvss:        0.0,
		Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
		Description: "    insecure_server.add_insecure_port('[::]:{}'.format(flags.port))",
	}

	assert.Equal(t, expectedIssue2, issues[1])

}
