package color

import (
	"errors"
	"strconv"
)

func ProcessRGB(s string) ([]int, error) {
	res := []int{}
	val := ""
	for _, char := range s {
		if char >= '0' && char <= '9' {
			val += string(char)
		} else if (char == ' ' || char == ',' || char == ')') && val != "" {
			x, _ := strconv.Atoi(val)
			if x < 0 || x > 255 {
				return []int{}, errors.New("one of the values is less than 0 or greater than 255")
			}
			res = append(res, x)
			val = ""
		}
	}
	if len(res) != 3 {
		return []int{}, errors.New("too many rgb int values passed")
	}

	return res, nil
}
