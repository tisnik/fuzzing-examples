package example4

import "fmt"

const maxWidth = 1024
const maxHeight = 1024

func TestedFunction(width int32, height int32) {
	if width < maxWidth && height < maxHeight {
		size := 4 * width * height
		bitmap := make([]byte, size)
		fmt.Println(len(bitmap))
	}
}
