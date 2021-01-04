// Package main implements a Dracon wrapper for MobSF, a mobile security
// framework (https://github.com/MobSF/Mobile-Security-Framework-MobSF). The
// wrapper handles the initialisation of MobSF, the identification of individual
// MobSF-compatible mobile app projects within the target code base, the
// compression and uploading of these projects to MobSF, and the retrieval of
// scan reports for processing by the Dracon MobSF producer.
package main

import (
	"archive/zip"
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	dtemplate "github.com/thought-machine/dracon/pkg/template"
)

const (
	MobSFBindHost = "127.0.0.1"
	MobSFBindPort = 8080
)

var MobSFAPIKey = generateAPIKey()

// CLI represents the command line options supported by this tool.
type CLI struct {
	InPath  string
	OutPath string
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

// parseCLIOptions returns a CLI struct representing the command line options
// that were passed to this tool.
func parseCLIOptions() *CLI {
	cli := new(CLI)

	flag.StringVar(
		&cli.InPath,
		"in",
		dtemplate.TemplateVars.ProducerSourcePath,
		"Path to directory containing source code to scan",
	)

	flag.StringVar(
		&cli.OutPath,
		"out",
		dtemplate.TemplateVars.ProducerToolOutPath,
		"Path to directory to which MobSF scan result reports should be written",
	)

	flag.Parse()

	return cli
}

// generateAPIKey generates an API key for the MobSF REST API. The generated key
// has low entropy, but MobSF listens locally and only runs for the duration of
// the code base scan, so this shouldn't be a problem.
func generateAPIKey() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	hash := sha256.Sum256([]byte(strconv.Itoa(r.Int())))
	return fmt.Sprintf("%x", hash)
}

// startMobSF spawns MobSF in a child process and returns the PID of that
// process.
func startMobSF() (int, error) {
	bindArg := fmt.Sprintf("--bind=%s:%d", MobSFBindHost, MobSFBindPort)

	mobSF := exec.Command(
		"/usr/local/bin/gunicorn",
		"--daemon", bindArg, "--workers=1", "--threads=10", "--timeout=1800",
		"MobSF.wsgi:application",
	)
	mobSF.Env = append(os.Environ(), "MOBSF_API_KEY=" + MobSFAPIKey)
	mobSF.Dir = "/root/Mobile-Security-Framework-MobSF"

	if err := mobSF.Run(); err != nil {
		return 0, err
	}

	return mobSF.Process.Pid, nil
}

// isAPIResponsive returns true if the MobSF REST API is responding to requests,
// or false if it is not.
func isAPIResponsive() bool {
	client := http.Client{Timeout: 1 * time.Second}

	resp, err := client.Get(fmt.Sprintf("http://%s:%d/api/v1", MobSFBindHost, MobSFBindPort))
	if err != nil {
		return false
	}

	return strings.Contains(resp.Header.Get("Access-Control-Allow-Headers"), "Authorization")
}

// findProjects searches a directory tree for Android and iOS projects, and
// returns a list of directories that represent the root directories for the
// discovered projects.
func findProjects(path string) ([]string, error) {
	if dir, err := os.Stat(path); err != nil || !dir.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", path)
	}

	projectDirs := []string{}
	err := filepath.Walk(path, func(f string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fi.IsDir() {
			return nil
		}

		isProjectDir := false

		if isAndroidEclipseProject(f) {
			log.Printf("Found Android Eclipse project at %s\n", f)
			isProjectDir = true
		} else if isAndroidStudioProject(f) {
			log.Printf("Found Android Studio project at %s\n", f)
			isProjectDir = true
		} else if isXcodeiOSProject(f) {
			log.Printf("Found Xcode iOS project at %s\n", f)
			isProjectDir = true
		}

		// If this directory is a supported project type, there's no need to
		// walk its contents
		if isProjectDir {
			projectDirs = append(projectDirs, f)
			return filepath.SkipDir
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to walk directory %s: %w", path, err)
	}

	return projectDirs, nil
}

// isAndroidEclipseProject returns true if the given path is the root directory
// of an Android Eclipse project, or false if not.
//
// The following conditions must hold for a directory structure to be considered
// an Android Eclipse project by MobSF:
// - the file AndroidManifest.xml must exist;
// - the directory src/ must exist.
func isAndroidEclipseProject(path string) bool {
	if f, err := os.Stat(path); err != nil || !f.IsDir() {
		return false
	}

	f := filepath.Join(path, "AndroidManifest.xml")
	fi, err := os.Stat(f)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Printf("Failed to stat %s: %v\n", f, err)
		}
		return false
	}
	if !fi.Mode().IsRegular() {
		return false
	}

	f = filepath.Join(path, "src")
	fi, err = os.Stat(f)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Printf("Failed to stat %s: %v\n", f, err)
		}
		return false
	}
	if !fi.IsDir() {
		return false
	}

	return true
}

// isAndroidStudioProject returns true if the given path is the root directory
// of an Android Studio project, or false if not.
//
// The following conditions must hold for a directory structure to be considered
// an Android Studio project by MobSF:
// - the file app/src/main/AndroidManifest.xml must exist;
// - the directory app/src/main/java/ must exist.
func isAndroidStudioProject(path string) bool {
	if f, err := os.Stat(path); err != nil || !f.IsDir() {
		return false
	}

	f := filepath.Join(path, "app", "src", "main", "AndroidManifest.xml")
	fi, err := os.Stat(f)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Printf("Failed to stat %s: %v\n", f, err)
		}
		return false
	}
	if !fi.Mode().IsRegular() {
		return false
	}

	f = filepath.Join(path, "app", "src", "main", "java")
	fi, err = os.Stat(f)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Printf("Failed to stat %s: %v\n", f, err)
		}
		return false
	}
	if !fi.IsDir() {
		return false
	}

	return true
}

// isXcodeiOSProject returns true if the given path is the root directory of an
// Xcode iOS project, or false if not.
//
// The following conditions must hold for a directory structure to be considered
// an Xcode iOS project by MobSF:
// - a directory with the suffix .xcodeproj must exist;
// - a file named Info.plist must exist somewhere in the directory structure.
func isXcodeiOSProject(path string) bool {
	if f, err := os.Stat(path); err != nil || !f.IsDir() {
		return false
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Printf("Failed to read directory %s: %v\n", path, err)
		return false
	}
	hasXcodeproj := false
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".xcodeproj") {
			hasXcodeproj = true
			break
		}
	}
	if !hasXcodeproj {
		return false
	}

	hasinfoPlist := false
	err = filepath.Walk(path, func(f string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if _, fc := filepath.Split(f); fc == "Info.plist" && fi.Mode().IsRegular() {
			hasinfoPlist = true
			return io.EOF  // Target file found - stop walking
		}

		return nil
	})
	if err != nil && err != io.EOF {
		log.Printf("Failed to walk directory %s: %v\n", path, err)
		return false
	}
	if !hasinfoPlist {
		return false
	}

	return true
}

// generateZipFile generates a zip file containing the files in the directory at
// the given path and writes the resulting zip file to the given io.Writer.
func generateZipFile(path string, dest io.Writer) error {
	if dir, err := os.Stat(path); err != nil || !dir.IsDir() {
		return fmt.Errorf("%s is not a directory", path)
	}

	zipWriter := zip.NewWriter(dest)
	err := filepath.Walk(path, func(f string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fi.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(path, f)
		if err != nil {
			return err
		}

		log.Printf("Adding: %s\n", relPath)

		zipMember, err := zipWriter.Create(relPath)
		if err != nil {
			return err
		}

		fh, err := os.Open(f)
		if err != nil {
			return err
		}

		if _, err = io.Copy(zipMember, fh); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	zipWriter.Close()
	if err != nil {
		return err
	}

	return nil
}

// uploadProjects uploads the project code bases in the given directories to
// MobSF, and returns a mapping from each project directory path to its
// corresponding MobSFFile representing the zip file containing that project
// code base in MobSF.
func uploadProjects(projectDirs []string) (map[string]*MobSFFile, error) {
	mobSFFiles := make(map[string]*MobSFFile, len(projectDirs))

	for _, dir := range projectDirs {
		// Generate zip file from project code base
		log.Printf("Generating zip file for project in %s\n", dir)
		body := new(bytes.Buffer)
		multipartWriter := multipart.NewWriter(body)
		zipName := fmt.Sprintf("%x.zip", sha256.Sum256([]byte(dir)))
		part, err := multipartWriter.CreateFormFile("file", zipName)
		if err != nil {
			return nil, fmt.Errorf("failed to open multipart writer: %w", err)
		}
		if err := generateZipFile(dir, part); err != nil {
			return nil, fmt.Errorf("failed to generate zip file: %w", err)
		}
		if err := multipartWriter.Close(); err != nil {
			return nil, fmt.Errorf("failed to close multipart writer: %w", err)
		}

		// Build upload request
		url := fmt.Sprintf("http://%s:%d/api/v1/upload", MobSFBindHost, MobSFBindPort)
		req, err := http.NewRequest("POST", url, body)
		if err != nil {
			return nil, fmt.Errorf("failed to create API request: %w", err)
		}
		req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
		req.Header.Set("Authorization", MobSFAPIKey)

		// Send upload request
		log.Println("Uploading zip file to MobSF")
		client := &http.Client{Timeout: 5 * time.Minute}
		resp, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("MobSF API request failed: %w", err)
		}
		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("MobSF API request failed: %s", resp.Status)
		}
		if !strings.HasPrefix(resp.Header.Get("Content-Type"), "application/json") {
			return nil, fmt.Errorf("MobSF API did not respond with JSON content")
		}

		// Parse response
		var file *MobSFFile
		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(&file); err != nil {
			return nil, fmt.Errorf("unable to parse JSON response from MobSF API: %w", err)
		}

		mobSFFiles[dir] = file
	}

	return mobSFFiles, nil
}

// scanProject orders MobSF to scan the given project code base, and returns a
// string containing a JSON-encoded scan report that can be parsed by the Dracon
// MobSF producer.
func scanProject(file *MobSFFile) (string, error) {
	// Build scan request
	url := fmt.Sprintf("http://%s:%d/api/v1/scan", MobSFBindHost, MobSFBindPort)
	req, err := http.NewRequest("POST", url, strings.NewReader(file.AsScanQuery()))
	if err != nil {
		return "", fmt.Errorf("failed to create API scan request: %w", err)
	}
	req.Header.Set("Authorization", MobSFAPIKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send scan request
	client := &http.Client{Timeout: 10 * time.Minute}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("MobSF API scan request failed: %w", err)
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("MobSF API scan request failed: %s", resp.Status)
	}
	if !strings.HasPrefix(resp.Header.Get("Content-Type"), "application/json") {
		return "", fmt.Errorf("MobSF API did not respond with JSON content on scan endpoint")
	}

	// Build JSON report request
	url = fmt.Sprintf("http://%s:%d/api/v1/report_json", MobSFBindHost, MobSFBindPort)
	req, err = http.NewRequest("POST", url, strings.NewReader(file.AsReportQuery()))
	if err != nil {
		return "", fmt.Errorf("failed to create API report request: %w", err)
	}
	req.Header.Set("Authorization", MobSFAPIKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send JSON report request
	client = &http.Client{Timeout: 30 * time.Second}
	resp, err = client.Do(req)
	if err != nil {
		return "", fmt.Errorf("MobSF API report request failed: %w", err)
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("MobSF API report request failed: %s", resp.Status)
	}
	if !strings.HasPrefix(resp.Header.Get("Content-Type"), "application/json") {
		return "", fmt.Errorf("MobSF API did not respond with JSON content on report endpoint")
	}

	// Return JSON report response (the MobSF producer will parse this)
	reportBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read MobSF API report response body: %w", err)
	}
	resp.Body.Close()
	return string(reportBytes), nil
}

func main() {
	cli := parseCLIOptions()

	mobSFPid, err := startMobSF()
	if err == nil {
		log.Printf("Started MobSF (PID: %d)\n", mobSFPid)
	} else {
		log.Fatalf("Failed to start MobSF: %v\n", err)
	}

	log.Println("Waiting for MobSF REST API to become responsive")
	apiResponsive := false
	for range [30]int{} {
		if isAPIResponsive() {
			apiResponsive = true
			break
		} else {
			time.Sleep(1 * time.Second)
		}
	}
	if !apiResponsive {
		log.Fatalln("MobSF REST API did not become responsive within 30 seconds; exiting")
	}

	log.Printf("Searching for project directories in %s\n", cli.InPath)
	projectDirs, err := findProjects(cli.InPath)
	if err != nil {
		log.Fatalf("Failed while searching for project directories: %v\n", err)
	}

	log.Println("Uploading project code bases to MobSF")
	files, err := uploadProjects(projectDirs)
	if err != nil {
		log.Fatalf("Failed to upload project code bases to MobSF: %v\n", err)
	}

	for dir, file := range files {
		log.Printf("Scanning project in %s\n", dir)

		report, err := scanProject(file)
		if err != nil {
			log.Fatalf("Failed to scan project: %v\n", err)
		}

		reportDir, err := filepath.Rel(cli.InPath, dir)
		if err != nil {
			log.Fatalf("Failed to derive output directory for scan report: %v\n", err)
		}
		reportDir = filepath.Join(cli.OutPath, reportDir)
		reportPath := filepath.Join(reportDir, "mobsf-scan.json")

		log.Printf("Writing scan report to %s\n", reportPath)
		if err := os.MkdirAll(reportDir, 0755); err != nil {
			log.Fatalf("Failed to create directory %s: %v\n", reportDir, err)
		}
		f, err := os.Create(reportPath)
		if err != nil {
			log.Fatalf("Failed to open %s for writing: %v\n", reportPath, err)
		}
		if _, err := f.WriteString(report); err != nil {
			log.Fatalf("Failed to write scan report to %s: %v\n", reportPath, err)
		}
		f.Close()
	}
}
