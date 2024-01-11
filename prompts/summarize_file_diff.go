package prompts

import (
	_ "embed"
	"text/template"
	"bytes"
)

//go:embed summarize_file_diff.tpl
var summarizeFileDiffTemplate string

var summarizeFileDiff = template.Must(template.New("summarize_file_diff").Parse(summarizeFileDiffTemplate))

func SummarizeFileDiff(fileDiffs string) string {
	fileDiffs = OmitLongLines(fileDiffs)

	buf := bytes.NewBuffer(nil)
	err := summarizeFileDiff.Execute(buf, map[string]any{
		"fileDiffs": fileDiffs,
	})
	if err != nil {
		panic(err)
	}
	return buf.String()
}
