package stringsext

import "strings"

// Text returns lines as multiline text
func Text(lines ...string) string {
	return strings.Join(lines, "\n")
}
