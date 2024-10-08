package main

import (
	"fmt"
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

func main() {
	router := http.NewServeMux()
	api := humago.New(router, huma.DefaultConfig("My API", "1.0.0"))

	user_openapi.Router(api)

	fmt.Printf("Server running; \033[34;1;4mhttp://localhost:%s/docs\033[0m\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
