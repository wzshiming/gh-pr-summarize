package prompts

import (
	_ "embed"
	"text/template"
	"bytes"
)

//go:embed summarize_discussion.tpl
var summarizeDiscussionTemplate string

var summarizeDiscussion = template.Must(template.New("summarize_discussion").Parse(summarizeDiscussionTemplate))

type Comment struct {
	User string
	Body string
}

func SummarizeDiscussion(title string, cs []Comment) string {
	buf := bytes.NewBuffer(nil)
	err := summarizeDiscussion.Execute(buf, map[string]any{
		"title":      title,
		"discussion": cs,
	})
	if err != nil {
		panic(err)
	}
	return buf.String()
}
