package routes

import (
	"go_server/app/components"
	"io"
	"net/http"
	"os"
)

func aboutPage(env string) string {
	header := components.HeaderMolecule(env, "About")
	content := `<h2>About</h2>`
	return components.SiteLayout(header, content)
}

func AboutRoute(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, aboutPage(os.Getenv("APP_ENV")))
}
