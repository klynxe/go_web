package auth

import (
	"main/ExtError"
	"main/models/auth"
	"main/models/signUp"
)

type Intr interface {
	StartService(config Config) (extErr *ExtError.Error)
	SignUp(su signUp.Request) (extErr *ExtError.Error)
	Login(au auth.Request) (status int, extErr *ExtError.Error)
}
