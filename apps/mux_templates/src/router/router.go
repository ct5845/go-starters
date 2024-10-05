package router

import "net/http"

func HandleRoutes() {

	StaticRouteHandler()

	http.HandleFunc("GET /about", AboutRoute)
	http.HandleFunc("GET /", HomeRoute)
}
