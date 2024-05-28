package color

import (
	c "asciicolor/colors"
)

func Painter(inputParts [][]TobeColored, fileData []string, color []int) string {
	res := ""

	// Loop over the 2D array of the inputParts to be colored
	for _, val := range inputParts {
		for i := 0; i < 8; i++ {
			for _, part := range val {
				if part.str == "" {
					res += "\n"
					continue
				}

				for _, char := range part.str {
					startIndex := GetStartingIndex(int(char))
					if startIndex >= 0 {
						// check if the string is not supposed to be colored
						if !part.colored {
							// Append the art generated art to the result variable
							res += fileData[startIndex+i]
						}
						// check if string passed is to be colored, call a coloring function
						//  to append the coloring escape codes plus string to be colored
						// and the reset code
						if part.colored {
							r := c.RGB_CONVERTER([]int(color), fileData[startIndex+i])
							// Append the colored art to the result variable
							res += r
						}
					}
				}
			}
			res += "\n"
		}
	}

	return res
}

func Painter2(inputParts []string, fileData []string, color []int) string {
	res := ""
	for _, part := range inputParts {
		if part == "" {
			res += "\n"
			continue
		}
		for i := 0; i < 8; i++ {
			for _, char := range part {
				startingIndex := GetStartingIndex(int(char))
				if startingIndex >= 0 {
					// color the whole string
					x := c.RGB_CONVERTER([]int(color), fileData[startingIndex+i])
					res += x
				}
			}
			res += "\n"
		}
	}
	return res
}
