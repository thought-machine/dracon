package types

// Position represents where in the file the finding is located
type Position struct {
	Character int `json:"character"`
	Line      int `"json:line"`
	Position  int `"json:position"`
}

// TSLintIssue represents a TSLint Result
type TSLintIssue struct {
	RuleName      string   `json:"ruleName"`
	Failure       string   `json:"failure"`
	Name          string   `json:"name"`
	StartPosition Position `json:"startPosition"`
	EndPosition   Position `json:"endPosition"`
}
