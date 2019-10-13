package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.Handle("/", http.FileServer(http.Dir("./../client/angular-cli/dist/angular-cli/")))

	http.ListenAndServe(":"+port, nil)

}
