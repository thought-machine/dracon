package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEsPush(t *testing.T) {
	want := "OK"
	scanUUID := "test-uuid"
	scanStartTime, _ := time.Parse("2006-01-02T15:04:05.000Z", "2020-04-13 11:51:53+01:00")
	esIn, _ := json.Marshal(&esDocument{
		ScanStartTime: scanStartTime,
		ScanID:        scanUUID,
		ToolName:      "es-unit-tests",
		Source:        "es-tests-source",
		Title:         "es-tests-title",
		Target:        "es-tests-target",
		Type:          "es-tests-type",
		Severity:      "es-tests-severity",
		CVSS:          "es-tests-cvss",
		Confidence:    "es-tests-confidence",
		Description:   "es-tests-description",
		FirstFound:    "es-tests-ff",
		FalsePositive: false,
	})

	esStub := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		assert.Equal(t, buf.String(), esIn)
		w.WriteHeader(200)
		w.Write([]byte(want))
	}))
	defer esStub.Close()

	if err := getESClient(); err != nil {
		log.Fatal(err)
	}

	esPush(esIn)

}
