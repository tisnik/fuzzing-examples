// +build gofuzz

package example1

func Fuzz(data []byte) int {
	TestedFunction(data)
	return 0
}
