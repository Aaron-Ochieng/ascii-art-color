package main

import (
	c "asciicolor/colors"
	utils "asciicolor/utils"
	"fmt"
	"os"
	"strings"
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

	//
	if len(args) == 2 {
		if strings.Contains(args[0], "--color=rgb") {
			err := utils.CheckRGBFormat(args[0])
			utils.ErrHandler(err)

			//
			rgbColors, e := utils.ProcessRGB(args[0])
			utils.ErrHandler(e)

			//
			paint := utils.Painter2(inputParts, fileData, rgbColors)
			fmt.Print(paint)
			return
		}
		ansii, err := c.AnsiiRGB(args[0])
		utils.ErrHandler(err)
		painted := utils.Painter2(inputParts, fileData, ansii)
		fmt.Print(painted)
		return
	}

	//
	if len(args) == 3 {
		var res [][]utils.TobeColored
		if strings.Contains(args[2], args[1]) {
			res = utils.SplitWord(args[1], inputParts)
		} else {
			res = utils.SplitRune(args[1], inputParts)
		}

		if strings.Contains(args[0], "--color=rgb") {

			utils.RGBArt(args[0], fileData, res)
			return
		}

		//
		ansii, err := c.AnsiiRGB(args[0])
		utils.ErrHandler(err)
		painted := utils.Painter(res, fileData, ansii)
		fmt.Print(painted)
		return
	}

}
