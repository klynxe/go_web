package serviceAuth

import (
	"main/ExtError"
	"main/interface/interfaceDb"
	"main/models/modelAuth"
	"main/services/serviceDb"
)

var service authService

type authService struct {
	db interfaceDb.DB
}

func StartService() (extErr *ExtError.Error) {
	service.db, extErr = serviceDb.GetDb()
	return
}

func SignUp(su modelAuth.SignUp) (extErr *ExtError.Error) {
	if service.db == interfaceDb.DB(nil) {
		return ExtError.New("SignUp, DB not open", 0)
	}
	extErr = service.db.SignUp(su.Login, su.Email, su.Password)
	return
}
