package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	connect := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connect)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.Handle("/", http.FileServer(http.Dir("./client/dist/angular-cli/")))

	http.ListenAndServe(":"+port, nil)

}
