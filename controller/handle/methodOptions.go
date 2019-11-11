package handle

import (
	"log"
	"net/http"
)

var HandleOptions = func(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	defer r.Body.Close()
	w.WriteHeader(http.StatusOK)
}
