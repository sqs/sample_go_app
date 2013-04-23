package main

import (
	"flag"
	"log"
	"net/http"
	"sample_go_app"
)

func main() {
	var addr = flag.String("http", ":8080", "HTTP service address")
	flag.Parse()

	http.HandleFunc("/", sample_go_app.Handler)
	log.Println("Starting HTTP server on", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
