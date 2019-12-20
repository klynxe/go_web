package main

import (
	"encoding/json"
	"fmt"
	"github.com/dpapathanasiou/go-recaptcha"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"main/controller/dicorator"
	"main/controller/handle"
	"main/services"
	"main/services/auth"
	"net/http"
	"os"
	"time"
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
	redisUrl := os.Getenv("REDIS_URL")
	options, err := redis.ParseURL(redisUrl)
	if err != nil {
		log.Fatal(err)
	}
	client := redis.NewClient(options)

	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong, err)

	err = client.Set("sessions:login:session_id", "token", 24*time.Hour).Err()
	if err != nil {
		panic(err)
	}

	connect := os.Getenv("DATABASE_URL")

	authConfig := auth.Config{DbConnect: connect,
		MailServer:   "*****.hoster.by",
		MailPort:     "25",
		MailPassword: "******",
		MailFrom:     "***@***.***"}

	authService := &auth.Auth{}
	err = authService.StartService(authConfig)
	if err != nil {
		log.Fatal(err)
	}
	services.SetAuther(authService)

	/*db.StartService(connect)
	auth.StartService()
	db, err := sql.Open("postgres", connect)
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

	router.HandleFunc("/sign-up", dicorator.AllowCors.Decor(handle.SignUp)).Methods(http.MethodPost)
	router.HandleFunc("/auth", dicorator.AllowCors.Decor(handle.Auth)).Methods(http.MethodPost)
	router.HandleFunc("/captcha", dicorator.AllowCors.Decor(handleCheckCaptcha)).Methods(http.MethodPost)

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(clientDir))))
	mux.CORSMethodMiddleware(router)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal("failed to start server", err)
	}
}
