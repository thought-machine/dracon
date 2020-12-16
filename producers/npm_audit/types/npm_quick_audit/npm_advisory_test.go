package npm_quick_audit

import (
	"net/http"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var httpmockFiles = map[string]string{
	"https://npmjs.com/advisories/666":  "producers/npm_audit/types/npm_quick_audit/npm_advisory_not_json",
	"https://npmjs.com/advisories/999":  "producers/npm_audit/types/npm_quick_audit/npm_advisory_no_advisorydata",
	"https://npmjs.com/advisories/1556": "producers/npm_audit/types/npm_quick_audit/npm_advisory_1556",
}

var advisoryAST = &AdvisoryData{
	CVEs:               []string{"CVE-2020-15168"},
	CWE:                "CWE-400",
	Overview:           "Node Fetch did not honor the size option after following a redirect, which means that when a content size was over the limit, a FetchError would never get thrown and the process would end without failure.\n\nFor most people, this fix will have a little or no impact. However, if you are relying on node-fetch to gate files above a size, the impact could be significant, for example: If you don't double-check the size of the data after fetch() has completed, your JS thread could get tied up doing work on a large file (DoS) and/or cost you money in computing.",
	PatchedVersions:    ">=2.6.1 <3.0.0-beta.1|| >= 3.0.0-beta.9",
	Recommendation:     "Upgrade to version 2.6.1 or 3.0.0-beta.9",
	References:         "",
	VulnerableVersions: "< 2.6.1 || >= 3.0.0-beta.1 < 3.0.0-beta.9",
}

func TestMain(m *testing.M) {
	httpmock.ActivateNonDefault(HTTPClient)

	os.Exit(m.Run())
}

func setup(t *testing.T) {
	for url, file := range httpmockFiles {
		httpmock.RegisterResponder("GET", url,
			func(req *http.Request) (*http.Response, error) {
				resp := httpmock.NewBytesResponse(200, httpmock.File(file).Bytes())
				resp.Header.Add("Content-Type", "application/json")
				return resp, nil
			},
		)
	}

	httpmock.RegisterResponder("GET", "https://npmjs.com/advisories/404",
		httpmock.NewStringResponder(404, ""))

	httpmock.RegisterNoResponder(httpmock.NewNotFoundResponder(t.Fatal))
}

func TestNewAdvisoryDataNotFound(t *testing.T) {
	setup(t)
	httpmock.ZeroCallCounters()

	advisory, err := NewAdvisoryData("https://npmjs.com/advisories/404")
	assert.Nil(t, advisory)
	assert.Error(t, err)

	assert.Equal(t, 1, httpmock.GetTotalCallCount())
}

func TestNewAdvisoryDataNotJSON(t *testing.T) {
	setup(t)
	httpmock.ZeroCallCounters()

	advisory, err := NewAdvisoryData("https://npmjs.com/advisories/666")
	assert.Nil(t, advisory)
	assert.Errorf(t, err, "npm Registry did not respond with JSON content")

	assert.Equal(t, 1, httpmock.GetTotalCallCount())
}

func TestNewAdvisoryDataValid(t *testing.T) {
	setup(t)
	httpmock.ZeroCallCounters()

	advisory, err := NewAdvisoryData("https://npmjs.com/advisories/1556")
	assert.NoError(t, err)
	assert.True(t, assert.ObjectsAreEqual(advisoryAST, advisory))

	assert.Equal(t, 1, httpmock.GetTotalCallCount())
}
