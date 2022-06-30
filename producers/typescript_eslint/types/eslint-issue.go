package types

// Position represents where in the file the finding is located

type Message struct {
	RuleID   string `json:"ruleId"`
	Severity int    `json:"severity"`
	Message  string `json:"message"`
	Line     int    `json:"line"`
	Column   int    `json:"column"`
}

// TSLintIssue represents a ESLint Result
type ESLintIssue struct {
	FilePath string    `json:"filePath"`
	Messages []Message `json:"messages`
}
