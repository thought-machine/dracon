package main

import (
	v1 "api/proto/v1"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	want             = "OK"
	info             = `{"Version":{"Number":"7.1.23"}}`
	scanUUID         = "test-uuid"
	scanStartTime, _ = time.Parse("2006-01-02T15:04:05.000Z", "2020-04-13 11:51:53+01:00")

	esIn, _ = json.Marshal(&esDocument{
		ScanStartTime:  scanStartTime,
		ScanID:         scanUUID,
		ToolName:       "es-unit-tests",
		Source:         "es-tests-source",
		Title:          "es-tests-title",
		Target:         "es-tests-target",
		Type:           "es-tests-type",
		Severity:       v1.Severity_SEVERITY_INFO,
		SeverityText:   "Info",
		CVSS:           0.01,
		Confidence:     v1.Confidence_CONFIDENCE_INFO,
		ConfidenceText: "Info",
		Description:    "es-tests-description",
		FirstFound:     scanStartTime,
		Count:          2,
		FalsePositive:  false,
	})
)

func TestEsPushBasicAuth(t *testing.T) {

	esIndex = "dracon-es-test"

	esStub := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		w.WriteHeader(200)
		if r.Method == "GET" {

			uname, pass, ok := r.BasicAuth()
			assert.Equal(t, uname, "foo")
			assert.Equal(t, pass, "bar")
			assert.Equal(t, ok, true)

			w.Write([]byte(info))
		} else if r.Method == "POST" {
			// assert non authed operation (write results to index)
			assert.Equal(t, buf.String(), string(esIn))
			assert.Equal(t, r.RequestURI, "/"+esIndex+"/_doc")

			uname, pass, ok := r.BasicAuth()
			assert.Equal(t, uname, "foo")
			assert.Equal(t, pass, "bar")
			assert.Equal(t, ok, true)

			w.Write([]byte(want))
		}

	}))
	defer esStub.Close()
	os.Setenv("ELASTICSEARCH_URL", esStub.URL)

	// basic auth ops
	basicAuthUser = "foo"
	basicAuthPass = "bar"
	assert.Nil(t, getESClient())
	esPush(esIn)
}
func TestEsPush(t *testing.T) {
	esStub := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		w.WriteHeader(200)
		if r.Method == "GET" {
			w.Write([]byte(info))
		} else if r.Method == "POST" {
			// assert non authed operation (write results to index)
			assert.Equal(t, buf.String(), string(esIn))
			assert.Equal(t, r.RequestURI, "/"+esIndex+"/_doc")
			w.Write([]byte(want))
		}
	}))
	defer esStub.Close()
	os.Setenv("ELASTICSEARCH_URL", esStub.URL)
	assert.Nil(t, getESClient())
	esPush(esIn)
}
