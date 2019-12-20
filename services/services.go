package services

import (
	"main/services/auth"
)

var srvAuth auth.Intr

func SetAuther(v auth.Intr) {
	srvAuth = v
}

func GetAuther() auth.Intr {
	return srvAuth
}
