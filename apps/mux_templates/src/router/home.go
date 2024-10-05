package router

import (
	"html/template"
	"io"
	"mux_templates/src/features/ui"
	"net/http"
)

func homePage() string {
	header := ui.HeaderMolecule("Home")
	content := template.HTML(`<h2 class="text-2xl">Home</h2>`)
	return ui.SiteLayout(header, content)
}

func HomeRoute(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, homePage())
}
