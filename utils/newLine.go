package color

import "strings"

func HandleNewLine(s string) ([]string, error) {
	s = strings.ReplaceAll(s, "\\n", "\n")
	spString := strings.Split(s, "\n")
	inputParts, err := CheckIllegalChar(spString)
	return inputParts, err
}
