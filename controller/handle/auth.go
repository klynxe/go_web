package handle

import (
	"encoding/json"
	"fmt"
	"github.com/dpapathanasiou/go-recaptcha"
	"log"
	"main/models/auth"
	"main/models/signUp"
	"main/services"
	"net/http"
)

var SignUp = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signUp request : %s", r)
	req := signUp.Request{}
	resp := signUp.Response{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	srvAuth := services.GetAuther()

	result, err := recaptcha.Confirm(r.Host, req.Token)
	if err != nil || result != true {
		log.Println(err)
		w.WriteHeader(http.StatusForbidden)
		resp = signUp.Response{signUp.CAPTCHA_FAILURE, "Captcha failure"}
	} else if err := srvAuth.SignUp(req); err != nil {
		fmt.Println("signUp error : %s", err)
		w.WriteHeader(http.StatusBadRequest)
		resp = signUp.Response{signUp.SIGN_UP, "Sign up in DB error"}
	} else {
		w.WriteHeader(http.StatusOK)
		resp = signUp.Response{signUp.OK, "Success"}
	}

	jr, err := json.Marshal(resp)
	if err != nil {
		log.Println("Error Marshal" + err.Error())
	}

	w.Write(jr)

}

var Auth = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("auth request : %s", r)
	req := auth.Request{}
	resp := auth.Response{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	srvAuth := services.GetAuther()

	result, err := recaptcha.Confirm(r.Host, req.CaptchaToken)
	if err != nil || result != true {
		log.Println(err)
		w.WriteHeader(http.StatusForbidden)
		resp = auth.Response{auth.CAPTCHA_FAILURE, "Captcha failure", 0}
	} else if status, err := srvAuth.Login(req); err != nil {
		fmt.Println("signUp error : %s", err)
		w.WriteHeader(http.StatusBadRequest)
		resp = auth.Response{auth.AUTH, "Auth in DB error", 0}
	} else {
		w.WriteHeader(http.StatusOK)
		resp = auth.Response{auth.OK, "Success", status}
	}

	jr, err := json.Marshal(resp)
	if err != nil {
		log.Println("Error Marshal " + err.Error())
	}

	w.Write(jr)

}
