package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Exclusions represents a list of MobSF static analysis scan rules whose
// findings should be ignored when scan reports are being processed by the tool.
// A rule is given by its ID (the value of the "id" key in the YAML files in the
// directories below), and must be prefixed with either "android." or "ios." as
// appropriate.
// - Android: https://github.com/MobSF/Mobile-Security-Framework-MobSF/tree/master/StaticAnalyzer/views/android/rules
// - iOS: https://github.com/MobSF/Mobile-Security-Framework-MobSF/tree/master/StaticAnalyzer/views/ios/rules
type Exclusions struct {
	All         []string
	PerPlatform map[string]map[string]bool
}

// String returns the Exclusions in its canonical string form (a comma-delimited
// list of values in the order in which they were added).
func (e *Exclusions) String() string {
	if e.All == nil {
		return ""
	}
	return strings.Join(e.All, ",")
}

// Set defines a value for the Exclusions, given a comma-delimited list of
// values as a string.
func (e *Exclusions) Set(value string) error {
	for _, id := range strings.Split(value, ",") {
		if found, _ := regexp.MatchString(`^(android|ios)\.`, id); !found {
			return fmt.Errorf("rule ID must begin with either 'android.' or 'ios.'")
		}

		e.All = append(e.All, id)

		split := strings.SplitN(id, ".", 2)
		platform := split[0]
		mobSFID := split[1]
		if _, found := e.PerPlatform[platform]; !found {
			e.PerPlatform[platform] = make(map[string]bool)
		}
		e.PerPlatform[platform][mobSFID] = true
	}

	return nil
}

// SetFor returns a map whose keys represent rule IDs that should be excluded
// when scanning projects for the given platform.
func (e *Exclusions) SetFor(platform string) map[string]bool {
	if _, found := e.PerPlatform[platform]; found {
		return e.PerPlatform[platform]
	}
	return map[string]bool{}
}

// CLI represents the command line options supported by this tool.
type CLI struct {
	InPath                 string
	OutPath                string
	CodeAnalysisExclusions Exclusions
}

// NewCLI creates and initialises a new CLI struct.
func NewCLI() *CLI {
	cli := new(CLI)

	cli.CodeAnalysisExclusions.PerPlatform = make(map[string]map[string]bool)

	return cli
}
