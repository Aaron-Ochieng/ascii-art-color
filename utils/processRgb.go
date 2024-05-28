package color

import (
	"errors"
	"regexp"
)

func CheckRGBFormat(s string) error {
	pattern := `^--color="rgb\(\d{1,3}, \d{1,3}, \d{1,3}\)"$|^--color=rgb\(\d{1,3}, \d{1,3}, \d{1,3}\)$`
	match, _ := regexp.MatchString(pattern, s)

	// checks if the string matches the input flag
	if !match {
		return errors.New("--color=... arguments does not match the required format")
	}
	return nil
}
