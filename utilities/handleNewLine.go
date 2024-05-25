package ascii

import "strings"

func HandleNewLine(s string) ([]string, error) {
	spString := strings.Split(s, "\n") // Split input by "\\n" to handle newline sequence
	// check for  illegal characters in the string
	inputParts, err := CheckIllegalChar(spString)
	return inputParts, err
}