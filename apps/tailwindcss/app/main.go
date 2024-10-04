package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"tailwindcss/app/routes"
	"tailwindcss/app/utils"
)

func main() {
	utils.Env("APP_ENV")
	http.HandleFunc("/", routes.HomeRoute)
	http.HandleFunc("/about", routes.AboutRoute)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server listening on %shttp://localhost:%s\n", utils.Colors["Yellow"], port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)

	if err != nil {
		log.Fatal(err)
	}
}
