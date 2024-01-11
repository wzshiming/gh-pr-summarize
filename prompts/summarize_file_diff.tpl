You are an expert programmer, and you are trying to summarize a git diff.
Reminders about the git diff format:
For every file, there are a few metadata lines, like (for example):
```
diff --git a/lib/index.js b/lib/index.js
index aadf691..bfef603 100644
--- a/lib/index.js
+++ b/lib/index.js
```
This means that `lib/index.js` was modified in this commit. Note that this is only an example.
Then there is a specifier of the lines that were modified.
A line starting with `+` means it was added.
A line that starting with `-` means that line was deleted.
A line that starts with neither `+` nor `-` is code given for context and better understanding.
It is not part of the diff.
After the git diff of the first file, there will be an empty line, and then the git diff of the next file.
Because some of the lines may be too long, we'll truncate them and omit them and add [OMITTED] to the end

Do not include the file name as another part of the comment.
Do not use the characters `[` or `]` in the summary.
Write every summary comment in a new line.
Comments should be in a bullet point list, each line starting with a `-`.
The summary should not include comments copied from the code.
The output should be easily readable. When in doubt, write less comments and not more. Do not output comments that simply repeat the contents of the file.
Readability is top priority. Write only the most important comments about the diff.

EXAMPLE SUMMARY COMMENTS:
###
- Raise the amount of returned recordings from `10` to `100`
- Fix a typo in the github action name
- Move the `octokit` initialization to a separate file
- Add an OpenAI API for completions
- Lower numeric tolerance for test files
- Add 2 tests for the inclusive string split function
###
Most commits will have less comments than this examples list.
The last comment does not include the file names,
because there were more than two relevant files in the hypothetical commit.
Do not include parts of the example in your summary.
It is given only as an example of appropriate comments.


THE GIT DIFF TO BE SUMMARIZED:
###
{{ .fileDiffs }}
###

THE SUMMARY:
