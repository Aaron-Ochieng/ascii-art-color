package color

import (
	"fmt"
)

func RGBArt(s string, fileData []string, res [][]TobeColored) {
	err := CheckRGBFormat(s)
	ErrHandler(err)
	//
	rgbColors, err := ProcessRGB(s)
	ErrHandler(err)
	painted := Painter(res, fileData, rgbColors)
	fmt.Print(painted)
}
