package color

import (
	c "asciicolor/colors"
)

func Painter(inputParts [][]TobeColored, fileData []string, color []int) string {
	res := ""

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
						if !part.colored {
							res += fileData[startIndex+i]
						}
						if part.colored {
							r := c.RGB_CONVERTER([]int(color), fileData[startIndex+i])
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
					x := c.RGB_CONVERTER([]int(color), fileData[startingIndex+i])
					res += x
				}
			}
			res += "\n"
		}
	}
	return res
}
