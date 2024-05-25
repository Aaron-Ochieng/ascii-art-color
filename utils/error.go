package color

import (
	"fmt"
	"os"
)

func ErrHandler(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
