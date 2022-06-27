package utils

import "strings"

func Indent(s string) string {
	lines := strings.Split(s, "\n")

	for i, line := range lines {
		lines[i] = "\t" + line
	}

	return strings.Join(lines, "\n")
}
