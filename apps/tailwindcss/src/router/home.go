package router

import (
	"io"
	"net/http"
	"tailwindcss/src/components"
)

func homePage() string {
	header := components.HeaderMolecule("Home")
	content := `<h2 class="text-2xl">Home</h2>`
	return components.SiteLayout(header, content)
}

func HomeRoute(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, homePage())
}
