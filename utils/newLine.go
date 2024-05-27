package color

import "strings"

func HandleNewLine(s string) ([]string, error) {
	spString := strings.Split(s, "\n")
	inputParts, err := CheckIllegalChar(spString)
	return inputParts, err
}
