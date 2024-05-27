package color

import "strings"

func HandleBackspace(s string) string {
	index := strings.Index(s, "\\b")

	if index == -1 {
		return s
	}
	if index == 0 {
		return HandleBackspace(s[1:])
	}
	return HandleBackspace(s[:index-1] + s[index+2:])
}
