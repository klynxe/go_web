package handle

import (
	"encoding/json"
	"fmt"
	"main/models/modelAuth"
	"main/services/serviceAuth"
	"net/http"
)

var HandleSignUp = func(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("signUp request : %s", r)
	su := modelAuth.SignUp{}
	if err := json.NewDecoder(r.Body).Decode(&su); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := serviceAuth.SignUp(su); err != nil {
		fmt.Printf("signUp error : %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
