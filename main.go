package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
)

func pageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	http.ServeFile(w, r, filepath.Join("./", r.URL.Path))
}

func staticHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	addr := flag.String("addr", ":4004", "HTTP Port Address")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", pageHandler)
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on port ", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
