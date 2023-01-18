package template

import (
	"testing"

	heredoc "github.com/makenowjust/heredoc/v2"
	"github.com/stretchr/testify/assert"
)

func TestSplitYAML(t *testing.T) {
	var tests = []struct {
		desc     string
		inBytes  []byte
		outBytes [][]byte
		err      error
	}{
		{
			desc: "a simple YAML file with no leading document separator",
			inBytes: []byte(heredoc.Doc(`
			a: Easy
			b:
			    c: 2
			    d:
			        - 3
			        - 4
			`)),
			outBytes: [][]byte{[]byte(heredoc.Doc(`
			a: Easy
			b:
			    c: 2
			    d:
			        - 3
			        - 4
			`))},
			err: nil,
		},
		{
			desc: "a simple YAML file with a leading document separator",
			inBytes: []byte(heredoc.Doc(`
			---
			a: Easy
			b:
			    c: 2
			    d:
			        - 3
			        - 4
			`)),
			outBytes: [][]byte{[]byte(heredoc.Doc(`
			a: Easy
			b:
			    c: 2
			    d:
			        - 3
			        - 4
			`))},
			err: nil,
		},
		{
			desc: "a YAML file with no leading document separator but a string value containing ---",
			inBytes: []byte(heredoc.Doc(`
			a: Easy
			b:
			    c: 2 ---
			    d:
			        - 3
			        - 4
			`)),
			outBytes: [][]byte{[]byte(heredoc.Doc(`
			a: Easy
			b:
			    c: 2 ---
			    d:
			        - 3
			        - 4
			`))},
			err: nil,
		},
		{
			desc: "a YAML file with both a leading document separator and a string value containing ---",
			inBytes: []byte(heredoc.Doc(`
			---
			a: Easy
			b:
			    c: 2 ---
			    d:
			        - 3
			        - 4
			`)),
			outBytes: [][]byte{[]byte(heredoc.Doc(`
			a: Easy
			b:
			    c: 2 ---
			    d:
			        - 3
			        - 4
			`))},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			outBytes, err := splitYAML(tt.inBytes)
			if tt.err == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
			assert.Equal(t, tt.outBytes, outBytes)
		})
	}
}
