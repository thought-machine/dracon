package main

import (
  "testing"
  "fmt"
	"github.com/stretchr/testify/assert"
	v1 "github.com/thought-machine/dracon/pkg/genproto/v1"
)


func TestReadXML(t *testing.T) {
  var bytes []byte
	issues := readXML(bytes)
	assert.Equal(t, len(issues),0)

	issues = readXML([]byte(exampleOutput))

  expectedIssues := make([]*v1.Issue,9)
  iss := &v1.Issue{
		Target:      "com/h3xstream/findsecbugs/injection/BasicInjectionDetector.java:48-188",
		Type:        "PATH_TRAVERSAL_IN",
		Title:       "Potential Path Traversal (file read)",
		Severity:    v1.Severity_SEVERITY_MEDIUM,
		Cvss:        0.0,
		Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
		Description: "This API (java/io/File.<init>(Ljava/lang/String;)V) reads a file whose location might be specified by user input",
  }
  expectedIssues[0] = iss
  methodIssue := iss
  methodIssue.Target = "com/h3xstream/findsecbugs/injection/BasicInjectionDetector.java:155-161"
  expectedIssues[1] = methodIssue

  sourceIssue0 := methodIssue
  sourceIssue0.Target = "com/h3xstream/findsecbugs/injection/BasicInjectionDetector.java:155-155"
  expectedIssues[2] = sourceIssue0

  sourceIssue1 := sourceIssue0
  sourceIssue1.Target = "com/h3xstream/findsecbugs/injection/BasicInjectionDetector.java:53-53"
  expectedIssues[3] = sourceIssue1

  iss = &v1.Issue{
		Target:      "com/h3xstream/findsecbugs/injection/InjectionPoint.java:26-52",
		Type:        "EI_EXPOSE_REP",
		Title:       "May expose internal representation by returning reference to mutable object",
		Severity:    v1.Severity_SEVERITY_LOW,
		Cvss:        0.0,
		Confidence:  v1.Confidence_CONFIDENCE_MEDIUM,
		Description: "com.h3xstream.findsecbugs.injection.InjectionPoint.getInjectableArguments() may expose internal representation by returning InjectionPoint.injectableArguments",
  }
  expectedIssues[5] = iss

  methodIssue = iss
  methodIssue.Target = "com/h3xstream/findsecbugs/injection/InjectionPoint.java:39-39"
  expectedIssues[6] = methodIssue

  fieldIssue := iss
  fieldIssue.Target = "com/h3xstream/findsecbugs/injection/InjectionPoint.java:"
  expectedIssues[7] = fieldIssue

  sourceIssue0 = methodIssue
  sourceIssue0.Target = "com/h3xstream/findsecbugs/injection/InjectionPoint.java:39-39"
  expectedIssues[8] = sourceIssue0
  
  found := 0
  for issu := range(issues){
    fmt.Println(issues[issu])
  }
  for i:=0; i<len(issues); i++{
    for j:=0; j<len(expectedIssues); j++{
      if expectedIssues[i].Target == issues[i].Target{
        found++
        // assert.EqualValues(t, expectedIssues[i].Target, issues[i].Target)
        assert.EqualValues(t, expectedIssues[i].Type, issues[i].Type)
        assert.EqualValues(t, expectedIssues[i].Title, issues[i].Title)
        assert.EqualValues(t, expectedIssues[i].Severity, issues[i].Severity)
        assert.EqualValues(t, expectedIssues[i].Cvss, issues[i].Cvss)
        assert.EqualValues(t, expectedIssues[i].Confidence, issues[i].Confidence)
        assert.EqualValues(t, expectedIssues[i].Description, issues[i].Description)
        break;
    }
  }}
  assert.Equal(t, len(expectedIssues),len(issues))
  assert.Equal(t,found,len(issues))
}
const exampleOutput = `
<?xml version="1.0" encoding="UTF-8"?>

<BugCollection version="3.1.12" sequence="0" timestamp="1586189624000" analysisTimestamp="1586189616607" release="">
  <Project projectName="">
    <Jar>/</Jar>
  </Project>
<BugInstance type="PATH_TRAVERSAL_IN" priority="2" rank="12" abbrev="SECPTI" category="SECURITY" instanceHash="ad9a7f908979ab4ffb356ccad009849e" instanceOccurrenceNum="0" instanceOccurrenceMax="0" cweid="22">
<ShortMessage>Potential Path Traversal (file read)</ShortMessage>
<LongMessage>This API (java/io/File.&lt;init&gt;(Ljava/lang/String;)V) reads a file whose location might be specified by user input</LongMessage>
<Class classname="com.h3xstream.findsecbugs.injection.BasicInjectionDetector" primary="true">
  <SourceLine classname="com.h3xstream.findsecbugs.injection.BasicInjectionDetector0" start="48" end="188" sourcefile="BasicInjectionDetector.java" 
  sourcepath="com/h3xstream/findsecbugs/injection/BasicInjectionDetector0.java">
	<Message>At BasicInjectionDetector.java:[lines 48-188]</Message>
  </SourceLine>
  <Message>In class com.h3xstream.findsecbugs.injection.BasicInjectionDetector</Message>
</Class>
<Method classname="com.h3xstream.findsecbugs.injection.BasicInjectionDetector" name="loadCustomSinks" signature="(Ljava/lang/String;Ljava/lang/String;)V" isStatic="false" primary="true">
  <SourceLine classname="com.h3xstream.findsecbugs.injection.BasicInjectionDetector" start="155" end="161" startBytecode="0" endBytecode="460" 
  sourcefile="BasicInjectionDetector.java" sourcepath="com/h3xstream/findsecbugs/injection/BasicInjectionDetector.java"/>
  <Message>In method com.h3xstream.findsecbugs.injection.BasicInjectionDetector.loadCustomSinks(String, String)</Message>
</Method>
<SourceLine classname="com.h3xstream.findsecbugs.injection.BasicInjectionDetector1" primary="true" start="155" end="155" startBytecode="5" endBytecode="5" 
sourcefile="BasicInjectionDetector.java" sourcepath="com/h3xstream/findsecbugs/injection/BasicInjectionDetector1.java">
  <Message>At BasicInjectionDetector.java:[line 155]</Message>
</SourceLine>
<String value="java/io/File.&lt;init&gt;(Ljava/lang/String;)V" role="Sink method">
  <Message>Sink method java/io/File.&lt;init&gt;(Ljava/lang/String;)V</Message>
</String>
<String value="0" role="Sink parameter">
  <Message>Sink parameter 0</Message>
</String>
<String value="com/h3xstream/findsecbugs/injection/BasicInjectionDetector.loadCustomSinks(Ljava/lang/String;Ljava/lang/String;)V parameter 1" role="Unknown source">
  <Message>Unknown source com/h3xstream/findsecbugs/injection/BasicInjectionDetector.loadCustomSinks(Ljava/lang/String;Ljava/lang/String;)V parameter 1</Message>
</String>
<SourceLine classname="com.h3xstream.findsecbugs.FindSecBugsGlobalConfig" start="53" end="53" startBytecode="1" 
endBytecode="1" sourcefile="FindSecBugsGlobalConfig.java" sourcepath="com/h3xstream/findsecbugs/FindSecBugsGlobalConfig.java">
  <Message>At FindSecBugsGlobalConfig.java:[line 53]</Message>
</SourceLine>
</BugInstance>
<BugInstance type="EI_EXPOSE_REP" priority="2" rank="18" abbrev="EI" category="MALICIOUS_CODE" instanceHash="ad7f9636b7e868a48901d62a558b8664" instanceOccurrenceNum="0" instanceOccurrenceMax="0" cweid="374">
<ShortMessage>May expose internal representation by returning reference to mutable object</ShortMessage>
<LongMessage>com.h3xstream.findsecbugs.injection.InjectionPoint.getInjectableArguments() may expose internal representation by returning InjectionPoint.injectableArguments</LongMessage>
<Class classname="com.h3xstream.findsecbugs.injection.InjectionPoint" primary="true">
  <SourceLine classname="com.h3xstream.findsecbugs.injection.InjectionPoint0" start="26" end="52" sourcefile="InjectionPoint.java" 
  sourcepath="com/h3xstream/findsecbugs/injection/InjectionPoint1.java">
	<Message>At InjectionPoint.java:[lines 26-52]</Message>
  </SourceLine>
  <Message>In class com.h3xstream.findsecbugs.injection.InjectionPoint</Message>
</Class>
<Method classname="com.h3xstream.findsecbugs.injection.InjectionPoint1" name="getInjectableArguments" signature="()[I" isStatic="false" primary="true">
  <SourceLine classname="com.h3xstream.findsecbugs.injection.InjectionPoint" start="39" end="39" startBytecode="0" endBytecode="46" sourcefile="InjectionPoint.java" sourcepath="com/h3xstream/findsecbugs/injection/InjectionPoint.java"/>
  <Message>In method com.h3xstream.findsecbugs.injection.InjectionPoint.getInjectableArguments()</Message>
</Method>
<Field classname="com.h3xstream.findsecbugs.injection.InjectionPoint" name="injectableArguments" signature="[I" isStatic="false" primary="true">
  <SourceLine classname="com.h3xstream.findsecbugs.injection.InjectionPoint2" sourcefile="InjectionPoint.java" 
  sourcepath="com/h3xstream/findsecbugs/injection/InjectionPoint2.java">
	<Message>In InjectionPoint.java</Message>
  </SourceLine>
  <Message>Field com.h3xstream.findsecbugs.injection.InjectionPoint.injectableArguments</Message>
</Field>
<SourceLine classname="com.h3xstream.findsecbugs.injection.InjectionPoint3" primary="true" start="39" end="39" startBytecode="4"
 endBytecode="4" sourcefile="InjectionPoint.java" sourcepath="com/h3xstream/findsecbugs/injection/InjectionPoint3.java">
  <Message>At InjectionPoint.java:[line 39]</Message>
</SourceLine>
</BugInstance>
</BugCollection>`
