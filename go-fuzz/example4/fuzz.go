// +build gofuzz

package example4

func Fuzz(data []byte) int {
	if len(data) >= 8 {
		width := int32(data[0]) + 256*int32(data[1]) + 65536*int32(data[2]) + 16777216*int32(data[3])
		height := int32(data[4]) + 256*int32(data[5]) + 65536*int32(data[6]) + 16777216*int32(data[7])
		TestedFunction(width, height)
	}
	return 0
}
