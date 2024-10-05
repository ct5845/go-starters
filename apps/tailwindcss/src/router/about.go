package router

import (
	"io"
	"net/http"
	"tailwindcss/src/components"
)

func aboutPage() string {
	header := components.HeaderMolecule("About")
	content := `<h2 class="text-2xl">About</h2>`
	return components.SiteLayout(header, content)
}

func AboutRoute(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, aboutPage())
}
