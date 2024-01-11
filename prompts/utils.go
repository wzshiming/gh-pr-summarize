package prompts

import (
	"strings"
)

func OmitLongLines(s string) string {
	var buf strings.Builder
	const (
		maxLineLength = 100
		omittedString = "[OMITTED]"
	)

	for _, line := range strings.Split(s, "\n") {
		if len(line) > maxLineLength+len(omittedString) {
			line = line[:maxLineLength] + omittedString
		}
		buf.WriteString(line)
		buf.WriteByte('\n')
	}
	return buf.String()
}
