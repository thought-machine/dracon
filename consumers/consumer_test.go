package consumers

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	v1 "github.com/thought-machine/dracon/pkg/genproto/v1"
	"github.com/thought-machine/dracon/pkg/putil"
)

func TestLoadToolResponse(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "dracon-test")
	assert.Nil(t, err)
	tmpFile, err := ioutil.TempFile(tmpDir, "dracon-test-*.pb")
	assert.Nil(t, err)
	defer os.Remove(tmpFile.Name())
	issues := []*v1.Issue{
		&v1.Issue{
			Target:      "/dracon/source/foobar",
			Title:       "/dracon/source/barfoo",
			Description: "/dracon/source/example.yaml",
		},
	}
	err = putil.WriteResults("test-tool", issues, tmpFile.Name())
	assert.Nil(t, err)

	log.Println(tmpDir)
	inResults = path.Dir(tmpDir)

	draconStartTime := time.Now().UTC()
	os.Setenv(EnvDraconStartTime, draconStartTime.Format(time.RFC3339))
	os.Setenv(EnvDraconScanID, "test-id")
	toolRes, err := LoadToolResponse()
	assert.Nil(t, err)
	log.Println(toolRes)

	assert.Equal(t, "test-tool", toolRes[0].GetToolName())
	assert.Equal(t, "test-id", toolRes[0].GetScanInfo().GetScanUuid())
}
