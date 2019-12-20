package db

import "main/ExtError"

type Intr interface {
	IsOpen() (open bool)
	Open(connect string) (extErr *ExtError.Error)
	CheckExistLogin(login string) (exist bool, extErr *ExtError.Error)

	Login(loginOrEmail, password string) (status int, extErr *ExtError.Error)
	SignUp(login, email, password string) (activationKey string, extErr *ExtError.Error)
	RemoveAccount(email string) (extErr *ExtError.Error)
	ConfirmAccount(login, activationKey string, status int) (statusOld int, extErr *ExtError.Error)
}
