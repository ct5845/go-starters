package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"tailwindcss/src/router"
	"tailwindcss/src/utils"
)

func main() {
	utils.LoadEnv()

	router.StaticRouteHandler()
	http.HandleFunc("/about", router.AboutRoute)
	http.HandleFunc("/", router.HomeRoute)

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
