package color

import "fmt"

func RGB_CONVERTER(color []int, s string) string {
	resetColor := "\x1b[0m"
	// converts int to unsigned integer8
	r, g, b := uint8(color[0]), uint8(color[1]), uint8(color[2])

	// front escape color code
	rgbColor := fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)

	// joining 1st ecape color code , string to be colored and the reset color
	coloredString := fmt.Sprintf("%s%s%s", rgbColor, s, resetColor)

	return coloredString
}
