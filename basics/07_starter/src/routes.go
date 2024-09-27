package main

import (
	"github.com/ct5845/htmx-starter/src/features/ui/pages"

	"net/http"
	"path/filepath"
)

func Routes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/static/", staticRoute)
	router.HandleFunc("/", defaultRoute)

	return router
}

func defaultRoute(w http.ResponseWriter, r *http.Request) {
	page := pages.HomePage()

	page.Render(r.Context(), w)
	// fmt.Fprintf(w, "Hello World")
}

func staticRoute(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path[len("/static/"):] // like doing r.URL.path.substring(9);
	fullPath := filepath.Join(".", "static", filePath)
	http.ServeFile(w, r, fullPath)
}
