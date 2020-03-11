package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	log.Println(request.URL)
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(writer, "can't read body", http.StatusBadRequest)
		return
	}
	log.Println(string(body))
	io.WriteString(writer, "Hello world!\n\n")
}

func main() {
	http.HandleFunc("/", mainEndpoint)
	http.ListenAndServe(":8080", nil)
}
