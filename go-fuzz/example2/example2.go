package example2

func TestedFunction(data []byte) bool {
	if len(data) >= 3 {
		if data[0] == 3 && data[1] == 2 && data[2] == 1 {
			return true
		}
	}
	return false
}
