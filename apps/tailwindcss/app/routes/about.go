package routes

import (
	"io"
	"net/http"
	"os"
	"tailwindcss/app/components"
)

func aboutPage(env string) string {
	header := components.HeaderMolecule(env, "About")
	content := `<h2>About</h2>`
	return components.SiteLayout(header, content)
}

func AboutRoute(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, aboutPage(os.Getenv("APP_ENV")))
}
