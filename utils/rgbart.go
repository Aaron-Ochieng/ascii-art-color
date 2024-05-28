package color

import (
	"fmt"
)

func RGBArt(s string, fileData []string, res [][]TobeColored) {
	// check if the rgb format is respected
	err := CheckRGBFormat(s)
	ErrHandler(err)

	// check the rgb values
	rgbColors, err := HandleRGB(s)
	ErrHandler(err)

	// Generate the colored art
	painted := Painter(res, fileData, rgbColors)
	fmt.Print(painted)
}
