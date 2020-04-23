package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	v1 "api/proto/v1"

	"github.com/thought-machine/dracon/producers"
)

type Jar struct {
	XMLName xml.Name `xml:"Jar"`
}

type Project struct {
	XMLName xml.Name `xml:"Project"`
	Jar     *Jar     `xml:"Jar"`
}

type Method struct {
	XMLName    xml.Name     `xml:"Method"`
	Classname  string       `xml:"classname,attr"`
	Name       string       `xml:"name,attr"`
	Signature  string       `xml:"signature,attr"`
	IsStatic   string       `xml:"isStatic,attr"`
	SourceLine []SourceLine `xml:"SourceLine"`
}
type SourceLine struct {
	XMLName       xml.Name `xml:"SourceLine"`
	Classname     string   `xml:"classname,attr"`
	Start         string   `xml:"start,attr"`
	End           string   `xml:"end,attr"`
	StartBytecode string   `xml:"startBytecode,attr"`
	EndBytecode   string   `xml:"endBytecode,attr"`
	Sourcefile    string   `xml:"sourcefile,attr"`
	Sourcepath    string   `xml:"sourcepath,attr"`
	Role          string   `xml:"role,attr"`
}
type Class struct {
	XMLName    xml.Name     `xml:"Class"`
	Classname  string       `xml:"classname,attr"`
	Role       string       `xml:"role,attr"`
	SourceLine []SourceLine `xml:"SourceLine"`
}
type Field struct {
	XMLName    xml.Name     `xml:"Field"`
	Classname  string       `xml:"classname,attr"`
	SourceLine []SourceLine `xml:"SourceLine"`
}
type BugInstance struct {
	XMLName      xml.Name     `xml:"BugInstance"`
	Class        []Class      `xml:"Class"`
	Method       []Method     `xml:"Method"`
	SourceLine   []SourceLine `xml:"SourceLine"`
	Field        []Field      `xml:"Field"`
	LongMessage  string       `xml:"LongMessage"`
	ShortMessage string       `xml:"ShortMessage"`
	Type         string       `xml:"type,attr"`
	Priority     string       `xml:"priority,attr"`
	Rank         string       `xml:"rank,attr"`
	Abbrev       string       `xml:"abbrev,attr"`
	Category     string       `xml:"category,attr"`
}
type BugCollection struct {
	XMLName     xml.Name      `xml:"BugCollection"`
	Project     *Project      `xml:"Project"`
	BugInstance []BugInstance `xml:"BugInstance"`
}

func loadXML(filename string) ([]byte, error) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer xmlFile.Close()
	return ioutil.ReadAll(xmlFile)

}
func readXML(xmlFile []byte) []*v1.Issue {
	/**
	Reads a file containing spotbugs XML results
	and converts the results in the "SECURITY" category
	into an array Dracon issues
	*/

	output := []*v1.Issue{}
	var bugs BugCollection
	if len(xmlFile) == 0 {
		return output
	}
	xml.Unmarshal(xmlFile, &bugs)
	for _, instance := range bugs.BugInstance {

		// parse standalone SourceLine elements
		for _, line := range instance.SourceLine {
			output = append(output, parseLine(instance, line))
		}
		// parse SourceLines in Field elements
		for _, field := range instance.Field {
			for _, line := range field.SourceLine {
				output = append(output, parseLine(instance, line))
			}
		}
		// parse SourceLines in Method elements
		for _, method := range instance.Method {
			for _, line := range method.SourceLine {
				output = append(output, parseLine(instance, line))
			}
		}
		//parse SourceLines in Class elements
		for _, cls := range instance.Class {
			for _, line := range cls.SourceLine {
				output = append(output, parseLine(instance, line))
			}
		}

	}
	return output
}
func parseLine(instance BugInstance, sourceLine SourceLine) *v1.Issue {
	return &v1.Issue{
		Target:      fmt.Sprintf("%s:%s-%s", sourceLine.Sourcepath, sourceLine.Start, sourceLine.End),
		Type:        instance.Type,
		Severity:    normalizeRank(instance.Rank),
		Cvss:        0.0,
		Confidence:  v1.Confidence(v1.Confidence_value[fmt.Sprintf("CONFIDENCE_%s", "MEDIUM")]),
		Description: instance.LongMessage,
		Title:       instance.ShortMessage,
	}
}
func normalizeRank(rank string) v1.Severity {
	/*
			Normalizes the rank according to the following table
			Scariest: ranked between 1 & 4.
		Scary: ranked between 5 & 9.
		Troubling: ranked between 10 & 14.
		Of concern: ranked between 15 & 20.
	*/
	intRank, err := strconv.ParseInt(rank, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	if intRank > 1 && intRank < 4 {
		return v1.Severity_SEVERITY_CRITICAL
	} else if intRank > 5 && intRank < 9 {
		return v1.Severity_SEVERITY_HIGH
	} else if intRank > 10 && intRank < 14 {
		return v1.Severity_SEVERITY_MEDIUM
	} else if intRank > 15 && intRank < 20 {
		return v1.Severity_SEVERITY_LOW
	}
	return v1.Severity_SEVERITY_INFO
}

func main() {
	if err := producers.ParseFlags(); err != nil {
		log.Fatal(err)
	}
	xmlByteVal, _ := loadXML(producers.InResults)
	issues := readXML(xmlByteVal)
	if err := producers.WriteDraconOut(
		"spotbugs",
		issues,
	); err != nil {
		log.Fatal(err)
	}
}
