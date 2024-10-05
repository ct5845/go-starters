package main

import (
	"fmt"
	"log"
	"mux_templates/src/router"
	"mux_templates/src/utils"
	"net/http"
	"os"
)

func main() {
	utils.LoadEnv()

	router.HandleRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server listening on \033[33;1;4mhttp://localhost:%s\033[0m\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)

	if err != nil {
		log.Fatal(err)
	}
}
