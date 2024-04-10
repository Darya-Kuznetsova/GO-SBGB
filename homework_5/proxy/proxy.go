package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

var (
	counter      int    = 0
	replica1Host string = "http://localhost:8081/create/"
	replica2Host string = "http://localhost:8082/create/"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", proxyHandler)

	http.ListenAndServe(":8083", mux)

}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	content, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	if counter == 0 {
		if _, err := http.Post(replica1Host, string(content), bytes.NewReader(content)); err != nil {
			log.Fatalln(err)
		}
		counter++
		return
	}
	if _, err := http.Post(replica2Host, string(content), bytes.NewReader(content)); err != nil {
		log.Fatalln(err)
	}
	counter--

}
