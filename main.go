package main

import (
	"fmt"
	"os"
	"strings"

	c "asciicolor/colors"
	utils "asciicolor/utils"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 || len(args) > 3 {
		fmt.Println("Usage: go run . --color=<color> <letters to be colored> something")
		return
	}

	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//
	fileData := utils.StringContain(string(file))
	if len(fileData) != 856 {
		fmt.Println("Error: >> Banner files  is empty with length of: ", len(file))
		return
	}

	var input string
	if len(args) == 2 {
		input = args[1]
	} else if len(args) == 3 {
		input = args[2]
	}
	if input == "" {
		return
	}
	if input == "\\n" {
		fmt.Println()
		return
	}
	input = utils.HandleBackspace(input)
	input = strings.ReplaceAll(string(input), "\\t", "   ")
	inputParts, err := utils.HandleNewLine(input)
	utils.ErrHandler(err)
	//

	// if the length of the arguments is 2
	if len(args) == 2 {
		// check if the flag contains rgb
		if strings.Contains(args[0], "--color=rgb") {
			// check the rgb format
			err := utils.CheckRGBFormat(args[0])
			utils.ErrHandler(err)

			// Extract the rgb values from the input flag
			rgbColors, e := utils.HandleRGB(args[0])
			utils.ErrHandler(e)

			// Add the color code and return the art
			paint := utils.Painter2(inputParts, fileData, rgbColors)

			// print the art
			fmt.Print(paint)
			return
		}
		// if the string.Contains doesnt match rgb... Execute the ansii color code
		ansii, err := c.AnsiiRGB(args[0])
		utils.ErrHandler(err)

		// Add the color code and return the art
		painted := utils.Painter2(inputParts, fileData, ansii)
		fmt.Print(painted)
		return
	}

	// if the length of the arguments == 3
	if len(args) == 3 {
		// initialize 2D array of strings supposed to be colored
		var res [][]utils.TobeColored

		// Checks if tobecolored substring exists in a string
		if strings.Contains(args[2], args[1]) {
			// If exists split the word according to the substring
			res = utils.SplitWord(args[1], inputParts)
		} else {
			// If it doenst, split the string according to runes
			res = utils.SplitRune(args[1], inputParts)
		}

		// If the flag contains rgb display the art according to rgb
		if strings.Contains(args[0], "--color=rgb") {
			utils.RGBArt(args[0], fileData, res)
			return // terminated the program
		}

		// Execute according to ansii and display its art
		ansii, err := c.AnsiiRGB(args[0])
		utils.ErrHandler(err)

		// Generate the colored ansii art
		painted := utils.Painter(res, fileData, ansii)
		fmt.Print(painted)
		return
	}
}
