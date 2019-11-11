package interfaceDb

import "main/ExtError"

type DB interface {
	IsOpen() (open bool)
	OpenDB(connect string) (extErr *ExtError.Error)
	Login(loginOrEmail, password string) (status int, extErr *ExtError.Error)
	CheckExistLogin(login string) (exist bool, extErr *ExtError.Error)
	SignUp(login, email, password string) (extErr *ExtError.Error)
}
