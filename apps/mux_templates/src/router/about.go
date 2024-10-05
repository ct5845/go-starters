package router

import (
	"html/template"
	"io"
	"mux_templates/src/features/ui"
	"net/http"
)

func aboutPage() string {
	header := ui.HeaderMolecule("About")
	content := template.HTML(`<h2 class="text-2xl">About</h2>`)
	return ui.SiteLayout(header, content)
}

func AboutRoute(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, aboutPage())
}
