package utils

import (
	"bytes"

	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

//TODO tests: process* count* get*
func TestPushMetrics(t *testing.T) {
	want := "OK"
	scanUUID := "test-uuid"
	scanStartTime, _ := time.Parse("2006-01-02T15:04:05.000Z", "2020-04-13 11:51:53+01:00")
	issuesNo := 1234
	slackIn := `{"text":"Dracon scan test-uuid started on 0001-01-01 00:00:00 +0000 UTC has been completed with 1234 issues\n"}`
	slackStub := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		assert.Equal(t, buf.String(), slackIn)
		w.WriteHeader(200)
		w.Write([]byte(want))
	}))
	defer slackStub.Close()
	PushMetrics(scanUUID, issuesNo, scanStartTime, slackStub.URL)

}

func TestPush(t *testing.T) {
	testMessage := "test Message"
	want := "OK"
	slackIn := `{"text":"` + testMessage + `"}`
	slackStub := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		assert.Equal(t, buf.String(), slackIn)
		w.WriteHeader(200)
		w.Write([]byte(want))
	}))
	defer slackStub.Close()

	PushMessage(testMessage, slackStub.URL)

}
