package producers

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
	v1 "github.com/thought-machine/dracon/pkg/genproto/v1"
)

func TestWriteDraconOut(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "dracon-test")
	assert.Nil(t, err)
	defer os.Remove(tmpFile.Name())
	outFile = tmpFile.Name()
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
}
