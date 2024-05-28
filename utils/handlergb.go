package color

import (
	"errors"
	"strconv"
)

func HandleRGB(s string) ([]int, error) {
	res := []int{}
	val := ""

	for _, char := range s {
		// check if the rune is in the range of 0 - 9
		if char >= '0' && char <= '9' {
			val += string(char)
			// check if char == space,comma,or closing braces and val is not empty
		} else if (char == ' ' || char == ',' || char == ')') && val != "" {
			x, _ := strconv.Atoi(val)
			// checks if the value is in the range of 0 - 255 and returns an error is not in the range of 0 - 255
			if x < 0 || x > 255 {
				return []int{}, errors.New("one of the values is less than 0 or greater than 255")
			}
			//
			res = append(res, x)
			val = ""
		}
	}

	// checks if the length of the values is exactly 3
	if len(res) != 3 {
		return []int{}, errors.New("too many rgb int values passed")
	}

	return res, nil
}
