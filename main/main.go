package main

import (
	"encoding/json"
	"github.com/dpapathanasiou/go-recaptcha"
	"github.com/gorilla/mux"
	"main/controller/dicorator"
	"main/controller/handle"
	"main/services/serviceAuth"
	"main/services/serviceDb"

	//"github.com/gomodule/redigo/redis"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

const (
	clientDir = "./client/angular_dist/"
)

type CheckCaptcha struct {
	Token string `json:"token"`
}

type ResultCheckCaptcha struct {
	Success bool `json:"success"`
}

func handleCheckCaptcha(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	checkCaptcha := CheckCaptcha{}
	// преобразуем json запрос в структуру
	err := decoder.Decode(&checkCaptcha)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result := false
	result, err = recaptcha.Confirm(r.Host, checkCaptcha.Token)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	jsonResult, _ := json.Marshal(ResultCheckCaptcha{result})
	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:4200")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResult)
}

func main() {

	connect := os.Getenv("DATABASE_URL")

	serviceDb.StartService(connect)
	serviceAuth.StartService()
	/*db, err := sql.Open("postgres", connect)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}*/

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	recaptcha.Init(os.Getenv("RECAPTCHA_PRIVATE_KEY"))

	router := mux.NewRouter()
	router.HandleFunc("/{*}", dicorator.AllowOptionsCors.Decor(handle.HandleOptions)).Methods(http.MethodOptions)

	router.HandleFunc("/sign-up", dicorator.AllowCors.Decor(handle.HandleSignUp)).Methods(http.MethodPost)
	router.HandleFunc("/captcha", dicorator.AllowCors.Decor(handleCheckCaptcha)).Methods(http.MethodPost)

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(clientDir))))
	mux.CORSMethodMiddleware(router)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal("failed to start server", err)
	}
}
