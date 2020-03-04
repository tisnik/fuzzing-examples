// +build gofuzz

package example5

func Fuzz(data []byte) int {
	return TestedFunction(data)
}
