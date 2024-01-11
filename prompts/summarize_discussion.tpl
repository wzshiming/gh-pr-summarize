You are an expert programmer, and you are trying to summarize discussions in GitHub PR comments.

Do not include the approved or rejected status of the PR in your summary.
Do not include the PR number in your summary.
Do not include the title of the PR in your summary.
Do not include comments in your summary.
Do not include URLs in your summary.
Ignore the automated comments from the CI system.
The summary should not include comments.

EXAMPLE SUMMARY:
###
- Add a new dependency to the build system
- Add a new test to the codebase
- Add a new feature to the codebase
- Fix a bug in a feature
- Fix a bug in the build system
- Fix flaky test
- Fix failing CI
###
Most commits will have less comments than this examples list.
Do not include parts of the example in your summary.
It is given only as an example of appropriate comments.

{{ with .title }}
THE TITLE:

{{ . }}

{{ end }}

THE DISCUSSION:

{{ range .discussion }}
[[[ BEGIN User @{{ .User }} ]]]
{{ .Body }}
[[[ END User @{{ .User }} ]]]
{{ end }}

THE SUMMARY:
