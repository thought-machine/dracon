package producers

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	v1 "github.com/thought-machine/dracon/api/proto/v1"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

type testJ struct {
	Foo string
}

func TestParseJSON(t *testing.T) {
	testJSON := `{"Foo":"bar"}`

	var inJSON testJ
	assert.Nil(t, ParseJSON(strings.NewReader(testJSON), &inJSON))
	assert.Equal(t, inJSON.Foo, "bar")
}

func TestWriteDraconOut(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "dracon-test")
	assert.Nil(t, err)
	defer os.Remove(tmpFile.Name())

	baseTime := time.Now().UTC()
	timestamp := baseTime.Format(time.RFC3339)
	os.Setenv(EnvDraconStartTime, timestamp)
	os.Setenv(EnvDraconScanID, "ab3d3290-cd9f-482c-97dc-ec48bdfcc4de")

	OutFile = tmpFile.Name()
	Append = false

	err = WriteDraconOut(
		"dracon-test",
		[]*v1.Issue{
			&v1.Issue{
				Target:      "/dracon/source/foobar",
				Title:       "/dracon/source/barfoo",
				Description: "/dracon/source/example.yaml",
			},
		},
	)
	assert.Nil(t, err)

	pBytes, err := ioutil.ReadFile(tmpFile.Name())
	res := v1.LaunchToolResponse{}
	err = proto.Unmarshal(pBytes, &res)
	assert.Nil(t, err)

	assert.Equal(t, "dracon-test", res.GetToolName())
	assert.Equal(t, "./foobar", res.GetIssues()[0].GetTarget())
	assert.Equal(t, "./barfoo", res.GetIssues()[0].GetTitle())
	assert.Equal(t, "./example.yaml", res.GetIssues()[0].GetDescription())
	assert.Equal(t, baseTime.Unix(), res.GetScanInfo().GetScanStartTime().GetSeconds())
	assert.Equal(t, "ab3d3290-cd9f-482c-97dc-ec48bdfcc4de", res.GetScanInfo().GetScanUuid())
}

func TestWriteDraconOutAppend(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "dracon-test")
	assert.Nil(t, err)
	defer os.Remove(tmpFile.Name())

	baseTime := time.Now().UTC()
	timestamp := baseTime.Format(time.RFC3339)
	os.Setenv(EnvDraconStartTime, timestamp)
	os.Setenv(EnvDraconScanID, "ab3d3290-cd9f-482c-97dc-ec48bdfcc4de")

	OutFile = tmpFile.Name()
	Append = true

	for _, i := range []int{0, 1, 2} {
		err = WriteDraconOut(
			"dracon-test",
			[]*v1.Issue{
				&v1.Issue{
					Target:      fmt.Sprintf("target%d", i),
					Title:       fmt.Sprintf("title%d", i),
					Description: fmt.Sprintf("desc%d", i),
				},
			},
		)
		assert.Nil(t, err)
	}

	pBytes, err := ioutil.ReadFile(tmpFile.Name())
	res := v1.LaunchToolResponse{}
	err = proto.Unmarshal(pBytes, &res)
	assert.Nil(t, err)

	assert.Equal(t, "dracon-test", res.GetToolName())
	assert.Equal(t, baseTime.Unix(), res.GetScanInfo().GetScanStartTime().GetSeconds())
	assert.Equal(t, "ab3d3290-cd9f-482c-97dc-ec48bdfcc4de", res.GetScanInfo().GetScanUuid())
	assert.Equal(t, 3, len(res.GetIssues()))

	for _, i := range []int{0, 1, 2} {
		assert.Equal(t, fmt.Sprintf("target%d", i), res.GetIssues()[i].GetTarget())
		assert.Equal(t, fmt.Sprintf("title%d", i), res.GetIssues()[i].GetTitle())
		assert.Equal(t, fmt.Sprintf("desc%d", i), res.GetIssues()[i].GetDescription())
	}
}
