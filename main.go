package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	scheme = "http"
)

func main() {
	addr := flag.String("addr", ":4004", "HTTP Port Address")
	flag.Parse()

	if *addr == ":443" {
		scheme = "https"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", pageHandler)
	mux.HandleFunc("/posts/", postHandler)
	mux.HandleFunc("/favicon.ico", faviconHandler)
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on port ", *addr)
	err := http.ListenAndServe(*addr, gzipHandler(redirectWWW(mux)))
	log.Fatal(err)
}
