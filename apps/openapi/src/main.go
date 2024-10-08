package main

import (
	"fmt"
	"net/http"
	"open_api/src/packages/user"
	"open_api/src/utils"

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

	user.Router(api)

	fmt.Printf("Server running; \033[34;1;4mhttp://localhost:%s/docs\033[0m\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
