package main

import (
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	fs := http.FileServer(http.Dir(filepath.Join("web", "html")))
	http.Handle("/", fs)

	log.Printf("Serving files from web/html on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
