// Package main implements a Dracon producer for MobSF, a mobile security
// framework (https://github.com/MobSF/Mobile-Security-Framework-MobSF). The
// producer acts as a wrapper around MobSF, handling the initialisation of the
// MobSF web server, the identification of individual MobSF-compatible mobile
// app projects within the target code base, the compression and uploading of
// these projects to MobSF, the retrieval of Android and iOS scan reports, and
// the conversion of scan reports into a Dracon Issues protobuf.
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
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/thought-machine/dracon/api/proto/v1"
	dtemplate "github.com/thought-machine/dracon/pkg/template"
	"github.com/thought-machine/dracon/producers"
	mreport "github.com/thought-machine/dracon/producers/mobsf/report"
	"github.com/thought-machine/dracon/producers/mobsf/report/android"
	"github.com/thought-machine/dracon/producers/mobsf/report/ios"
)

const (
	MobSFBindHost = "127.0.0.1"
	MobSFBindPort = 8080
)

var MobSFAPIKey = generateAPIKey()

// parseCLIOptions returns a CLI struct representing the command line options
// that were passed to this tool.
func parseCLIOptions() *CLI {
	cli := NewCLI()

	flag.StringVar(
		&cli.InPath,
		"in",
		dtemplate.TemplateVars.ProducerSourcePath,
		"Path to directory containing source code to scan",
	)

	flag.StringVar(
		&cli.OutPath,
		"out",
		dtemplate.TemplateVars.ProducerOutPath,
		"Path to which Dracon Issues protobuf should be written",
	)

	flag.Var(
		&cli.CodeAnalysisExclusions,
		"exclude",
		"Comma-delimited list of static analysis rule IDs to ignore",
	)

	flag.Parse()

	// So producers.WriteDraconOut knows where to write to:
	producers.OutFile = cli.OutPath

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
func findProjects(path string) ([]*Project, error) {
	if dir, err := os.Stat(path); err != nil || !dir.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", path)
	}

	projects := make([]*Project, 0)
	err := filepath.Walk(path, func(f string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fi.IsDir() {
			return nil
		}

		var projectType *ProjectType

		if isAndroidEclipseProject(f) {
			projectType = new(ProjectType)
			*projectType = AndroidEclipse
		} else if isAndroidStudioProject(f) {
			projectType = new(ProjectType)
			*projectType = AndroidStudio
		} else if isXcodeiOSProject(f) {
			projectType = new(ProjectType)
			*projectType = XcodeIos
		}

		if projectType != nil {
			log.Printf("Found %s project at %s\n", *projectType, f)
			projects = append(projects, &Project{
				RootDir: f,
				Type:    *projectType,
			})

			// If this directory is a supported project type, there's no need to
			// walk its contents
			return filepath.SkipDir
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to walk directory %s: %w", path, err)
	}

	return projects, nil
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
func uploadProjects(projects []*Project) error {
	for _, p := range projects {
		// Generate zip file from project code base
		log.Printf("Generating zip file for project in %s\n", p.RootDir)
		body := new(bytes.Buffer)
		multipartWriter := multipart.NewWriter(body)
		zipName := fmt.Sprintf("%x.zip", sha256.Sum256([]byte(p.RootDir)))
		part, err := multipartWriter.CreateFormFile("file", zipName)
		if err != nil {
			return fmt.Errorf("failed to open multipart writer: %w", err)
		}
		if err := generateZipFile(p.RootDir, part); err != nil {
			return fmt.Errorf("failed to generate zip file: %w", err)
		}
		if err := multipartWriter.Close(); err != nil {
			return fmt.Errorf("failed to close multipart writer: %w", err)
		}

		// Build upload request
		url := fmt.Sprintf("http://%s:%d/api/v1/upload", MobSFBindHost, MobSFBindPort)
		req, err := http.NewRequest("POST", url, body)
		if err != nil {
			return fmt.Errorf("failed to create API request: %w", err)
		}
		req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
		req.Header.Set("Authorization", MobSFAPIKey)

		// Send upload request
		log.Println("Uploading zip file to MobSF")
		client := &http.Client{Timeout: 5 * time.Minute}
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("MobSF API request failed: %w", err)
		}
		if resp.StatusCode != 200 {
			return fmt.Errorf("MobSF API request failed: %s", resp.Status)
		}
		if !strings.HasPrefix(resp.Header.Get("Content-Type"), "application/json") {
			return fmt.Errorf("MobSF API did not respond with JSON content")
		}

		// Parse response
		var file *MobSFFile
		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(&file); err != nil {
			return fmt.Errorf("unable to parse JSON response from MobSF API: %w", err)
		}

		p.Upload = file
	}

	return nil
}

// scanProject orders MobSF to scan the given project code base, ignoring the
// rule IDs given by exclusions, and returns a (partial) scan report.
func scanProject(project *Project, exclusions Exclusions) (mreport.Report, error) {
	// Build scan request
	url := fmt.Sprintf("http://%s:%d/api/v1/scan", MobSFBindHost, MobSFBindPort)
	req, err := http.NewRequest("POST", url, strings.NewReader(project.Upload.AsScanQuery()))
	if err != nil {
		return nil, fmt.Errorf("failed to create API scan request: %w", err)
	}
	req.Header.Set("Authorization", MobSFAPIKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send scan request
	client := &http.Client{Timeout: 10 * time.Minute}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("MobSF API scan request failed: %w", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("MobSF API scan request failed: %s", resp.Status)
	}
	if !strings.HasPrefix(resp.Header.Get("Content-Type"), "application/json") {
		return nil, fmt.Errorf("MobSF API did not respond with JSON content on scan endpoint")
	}

	// Build JSON report request
	url = fmt.Sprintf("http://%s:%d/api/v1/report_json", MobSFBindHost, MobSFBindPort)
	req, err = http.NewRequest("POST", url, strings.NewReader(project.Upload.AsReportQuery()))
	if err != nil {
		return nil, fmt.Errorf("failed to create API report request: %w", err)
	}
	req.Header.Set("Authorization", MobSFAPIKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send JSON report request
	client = &http.Client{Timeout: 30 * time.Second}
	resp, err = client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("MobSF API report request failed: %w", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("MobSF API report request failed: %s", resp.Status)
	}
	if !strings.HasPrefix(resp.Header.Get("Content-Type"), "application/json") {
		return nil, fmt.Errorf("MobSF API did not respond with JSON content on report endpoint")
	}

	// Parse response body as scan report
	reportBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read MobSF API report response body: %w", err)
	}
	resp.Body.Close()
	var report mreport.Report
	switch project.Type {
	case AndroidEclipse, AndroidStudio:
        report, err = android.NewReport(reportBytes, exclusions.SetFor("android"))
	case XcodeIos:
		report, err = ios.NewReport(reportBytes, exclusions.SetFor("ios"))
	default:
		return nil, fmt.Errorf("no report parser for this project type")
	}
	if err != nil {
		return nil, fmt.Errorf("error while parsing report: %w", err)
	}
	report.SetRootDir(project.RootDir)

	return report, nil
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
	projects, err := findProjects(cli.InPath)
	if err != nil {
		log.Fatalf("Failed while searching for project directories: %v\n", err)
	}

	log.Println("Uploading project code bases to MobSF")
	err = uploadProjects(projects)
	if err != nil {
		log.Fatalf("Failed to upload project code bases to MobSF: %v\n", err)
	}

	issues := make([]*v1.Issue, 0)
	for _, p := range projects {
		log.Printf("Scanning project in %s\n", p.RootDir)

		report, err := scanProject(p, cli.CodeAnalysisExclusions)
		if err != nil {
			log.Fatalf("Failed to scan project: %v\n", err)
		}

		reportIssues := report.AsIssues()
		log.Printf("Issues reported: %d\n", len(reportIssues))
		issues = append(issues, reportIssues...)
	}

	log.Printf("Writing Dracon Issues protobuf to %s\n", cli.OutPath)
	producers.WriteDraconOut("mobsf", issues)
}
