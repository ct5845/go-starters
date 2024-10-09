package main

import (
	"fmt"
	"insightful/src/group/group_api"
	"insightful/src/user/user_openapi"
	"insightful/src/utils"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
)

var port string

func init() {
	env := utils.LoadEnv()

	port = env.Port
	if port == "" {
		port = "8888"
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // You can specify your allowed origin here, e.g., "http://example.com"
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Request-ID")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	api := humago.New(mux, huma.DefaultConfig("My API", "1.0.0"))

	user_openapi.Router(api)
	group_api.Router(api)

	muxWithCors := corsMiddleware(mux)

	fmt.Printf("Server running; \033[34;1;4mhttp://localhost:%s/docs\033[0m\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), muxWithCors)
}
