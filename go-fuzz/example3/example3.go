package example3

import "fmt"

func TestedFunction(data []byte) {
	if len(data) >= 3 {
		if data[0] == 'r' && data[1] == 'o' && data[2] == 'o' && data[3] == 't' {
			fmt.Println("Spravny vstup")
		}
	}
}
