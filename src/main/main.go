package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {

	connect := os.Getenv("DATABASE_URL")
	if connect == "" {
		connect = "dbname=" + "klynxe" + " user=" + "egor" +
			" password=" + "12345" + " host=" + "localhost" + " port=" + "5432" + " sslmode=" + "disable"
	}
	fmt.Println("%v", t.T)
	_, err := sql.Open("postgres", connect)
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.Handle("/", http.FileServer(http.Dir("./src/client/angular-cli/dist/angular-cli/")))

	http.ListenAndServe(":"+port, nil)

}
