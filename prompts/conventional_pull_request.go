package prompts

import (
	_ "embed"
	"text/template"
	"bytes"
)

//go:embed conventional_pull_request.tpl
var conventionalPullRequestTemplate string

var conventionalPullRequest = template.Must(template.New("conventional_pull_request").Parse(conventionalPullRequestTemplate))

func ConventionalPullRequest(title string, cs []Comment, fileDiffs string) string {
	fileDiffs = OmitLongLines(fileDiffs)

	buf := bytes.NewBuffer(nil)
	err := conventionalPullRequest.Execute(buf, map[string]any{
		"title":      title,
		"discussion": cs,
		"fileDiffs":  fileDiffs,
	})
	if err != nil {
		panic(err)
	}
	return buf.String()
}

type Comment struct {
	User string
	Body string
}
