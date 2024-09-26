package main

import (
	"06_tailwind/libs/ui"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/a-h/templ"
)

func serveStatic(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path[len("/static/"):] // like doing r.URL.path.substring(9);
	fullPath := filepath.Join(".", "static", filePath)
	http.ServeFile(w, r, fullPath)
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/static/", serveStatic);
	router.Handle("/", templ.Handler(ui.Hello("templ")));

	fmt.Println("Listening on :8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}