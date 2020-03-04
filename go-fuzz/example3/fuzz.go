// +build gofuzz

package example3

func Fuzz(data []byte) int {
	TestedFunction(data)
	return 0
}
