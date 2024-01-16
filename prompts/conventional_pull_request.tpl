Your task is to categorize a code change discussed in a pull request.
After examining all changed files, the large ones were omitted from the diff summary, so you need to label it based on the available discussions and remaining file diffs.

Please use the following labels to categorize:

RELEASE: Involves changes related to the release of a new version of the project, often for version number updates or release notes.
BUILD: Involves changes impacting the build system or development environment.
CHORE: Related to minor routine tasks or repository maintenance activities, such as updating copyrights or repository settings.
CI: Pertains to changes made in the Continuous Integration (CI) configuration files and scripts used in the project development process.
DOCS: Represents non-code changes like editing or adding new documentation to improve project comprehension.
TYPO: Corrects typographical errors present in the code or documentation to enhance readability and understanding.
TEST: Associated with designing new tests for untested code, or correcting existing tests to improve codebase reliability.
PERFORMANCE: Centers around modifications aimed at improving the performance of the code, often for efficiency or scalability reasons.
REFACTOR: Involves changing the structure or coding architecture without altering the code behavior or adding a new feature, often for readability or other quality improvements.
STYLE: Consists of changes that don't alter the logic or function of the code (e.g., indentation, white-space, formatting, missing semi-colons, etc.), often to standardize or improve readability.
BUMP: Deals with updates or upgrades to libraries and dependencies to get benefits from their latest versions or fix security vulnerabilities.
FIX: Contains modifications that patch a bug or error existing in the codebase, often to improve functionality and reliability, not including tests, and typo fixes.
FEATURE: Introduces new functionalities like a new API, new configuration fields, new flags for command-line options, or enhancements to existing functionality, not including refactors and tests. This is often for expansion or improvement of the project's capabilities.
UNKNOWN: Should be chosen if it is not clear which label is appropriate given the current information.

Following are the summaries of the discussions and files:

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

THE FILE DIFF:

###
{{ .fileDiffs }}
###

Identify suitable labels from the given list.
You can choose multiple labels, each should be followed by a brief reason for the choice.

Your response should follow the format `LABEL: brief reason`.
For example, `FIX: Addressed a bug`. Short reasons are kept within 10 words.
If there are multiple labels, simply start a new `LABEL: brief reason` on a new line.
