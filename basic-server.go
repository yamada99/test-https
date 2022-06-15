package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "Hello world")
	})
	log.Println("start http listening :18888")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
