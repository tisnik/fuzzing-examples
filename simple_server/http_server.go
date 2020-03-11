package main

import (
	"io"
	"log"
	"net/http"
)

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	log.Println(request.URL)
	io.WriteString(writer, "Hello world!\n\n")
}

func main() {
	http.HandleFunc("/", mainEndpoint)
	http.ListenAndServe(":8080", nil)
}
