package router

import (
	"io"
	"mux_templates/src/features/about"
	"net/http"
)

func AboutRoute(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, about.Page())
}
