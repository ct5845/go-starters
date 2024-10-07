package router

import (
	"io"
	"mux_templates/src/features/home"
	"net/http"
)

func HomeRoute(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, home.Page())
}
