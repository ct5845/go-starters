package router

import (
	"net/http"
	"path/filepath"
	"tailwindcss/src/router/middleware"
	"tailwindcss/src/utils"
)

func StaticRouteHandler() {
	staticPath := "/static/"
	staticDir := filepath.Join(utils.AppDir(), "/static")
	staticFileHandler := http.StripPrefix(staticPath, http.FileServer(http.Dir(staticDir)))

	http.Handle(staticPath, middleware.GZip(staticFileHandler))
}
