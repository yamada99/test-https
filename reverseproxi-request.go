package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	director := func(request *http.Request) {
		request.URL.Scheme = "http"
		request.URL.Host = ":9001"
		request.URL.Path = "/new-path"
		request.Header.Add("New-Header", "From Proxy")
		request.Body = ioutil.NopCloser(strings.NewReader("body from proxy"))
	}
	rp := &httputil.ReverseProxy{
		Director: director,
	}
	server := http.Server{
		Addr:    "127.0.0.1:9000",
		Handler: rp,
	}
	log.Println("Start Listening at :9000")
	log.Fatalln(server.ListenAndServe())

}
