package routes

import (
	"go_server_docker/app/components"
	"io"
	"net/http"
	"os"
)

func homePage(env string) string {
	header := components.HeaderMolecule(env, "Home")
	content := `<h2>Home</h2>`
	return components.SiteLayout(header, content)
}

func HomeRoute(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, homePage(os.Getenv("APP_ENV")))
}
