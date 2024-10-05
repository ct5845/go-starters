package router

import (
	"mux_templates/src/router/middleware"
	"mux_templates/src/utils"
	"net/http"
	"path/filepath"
)

func StaticRouteHandler() {
	staticPath := "/static/"
	staticDir := filepath.Join(utils.AppDir(), "/static")
	staticFileHandler := http.StripPrefix(staticPath, http.FileServer(http.Dir(staticDir)))

	http.Handle("GET "+staticPath, middleware.GZip(staticFileHandler))
}
