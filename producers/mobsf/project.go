package main

import (
	"net/url"
)

// ProjectType represents a particular type of mobile app project supported by
// MobSF.
type ProjectType string

const (
	AndroidEclipse ProjectType = "Android Eclipse"
	AndroidStudio              = "Android Studio"
	XcodeIos                   = "Xcode iOS"
)

// Project represents a particular project somewhere in the target code base.
type Project struct {
	RootDir string
	Type    ProjectType
	Upload  *MobSFFile
}

// MobSFFile represents a file stored in MobSF. This is typically a project code
// base that has been uploaded to MobSF via the REST API or web interface.
type MobSFFile struct {
	FileName string `json:"file_name"`
	Hash     string `json:"hash"`
	ScanType string `json:"scan_type"`
}

// AsScanQuery returns a string representation of the MobSFFile that identifies
// the corresponding server-side file as part of a request to MobSF's scan
// endpoint.
func (m *MobSFFile) AsScanQuery() string {
	v := url.Values{}

	v.Add("file_name", m.FileName)
	v.Add("hash", m.Hash)
	v.Add("scan_type", m.ScanType)

	return v.Encode()
}

// AsReportQuery returns a string representation of the MobSFFile that
// identifies the corresponding server-side file as part of a request to any of
// MobSF's report generation endpoints.
func (m *MobSFFile) AsReportQuery() string {
	v := url.Values{}

	v.Add("hash", m.Hash)

	return v.Encode()
}
