package example5

import (
	"bytes"
	"image/gif"
)

func TestedFunction(data []byte) int {
	img, err := gif.Decode(bytes.NewReader(data))
	// chyba nastat muze - na vstupu jsou nahodna data
	if err != nil {
		// ovsem img by melo byt rovno nil
		if img != nil {
			panic("img != nil on error")
		}
		// jinak ok
		return 0
	}
	// pokus o zpetne vytvoreni obrazku
	// pro vsechny vstupy, ktere probehly "korektne"
	var w bytes.Buffer
	err = gif.Encode(&w, img, nil)
	if err != nil {
		panic(err)
	}
	return 1
}
