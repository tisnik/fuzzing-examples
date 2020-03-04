// +build gofuzz

package example2

func Fuzz(data []byte) int {
	if TestedFunction(data) {
		panic("wrong input")
	}
	return 0
}
