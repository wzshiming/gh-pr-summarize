package prompts

import (
	_ "embed"
	"text/template"
	"bytes"
)

//go:embed conventional_pull_request.tpl
var conventionalPullRequestTemplate string

var conventionalPullRequest = template.Must(template.New("conventional_pull_request").Parse(conventionalPullRequestTemplate))

func ConventionalPullRequest(summaryDiscussion, summaryPoints string) string {
	buf := bytes.NewBuffer(nil)
	err := conventionalPullRequest.Execute(buf, map[string]any{
		"summaryDiscussion": summaryDiscussion,
		"summaryPoints":     summaryPoints,
	})
	if err != nil {
		panic(err)
	}
	return buf.String()
}
