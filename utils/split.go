package color

import "strings"

func SplitWord(s1 string, s []string) [][]TobeColored {
	var arr [][]TobeColored

	// loop over the  slice of strings
	for _, s2 := range s {
		// split the string according to s1
		split := strings.Split(s2, s1)
		result := make([]TobeColored, 0, len(split)*2-1)


		// Given pinghellpot, it splits to ->["ping","pot"]
		for i, s := range split {
			if i != 0 {
				result = append(result, TobeColored{s1, true})
			}
			if s != "" {
				result = append(result, TobeColored{s, false})
			}
		}
		arr = append(arr, result)
	}

	return arr
}

func SplitRune(a string, b []string) [][]TobeColored {
	var arr [][]TobeColored
	for _, val := range b {
		var result []TobeColored
		temp := ""
		for _, runeValue := range val {
			if strings.ContainsRune(a, runeValue) {
				if temp != "" {
					result = append(result, TobeColored{str: temp, colored: false})
					temp = ""
				}
				result = append(result, TobeColored{str: string(runeValue), colored: true})
			} else {
				temp += string(runeValue)
			}
		}
		if temp != "" {
			result = append(result, TobeColored{str: temp, colored: false})
		}
		arr = append(arr, result)
	}

	return arr
}
