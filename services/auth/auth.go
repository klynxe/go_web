package auth

import (
	"encoding/base64"
	"fmt"
	"main/ExtError"
	"main/models/auth"
	"main/models/signUp"
	"main/resource"
	"main/services/auth/cash"
	"main/services/auth/cash/dbRedis"
	"main/services/auth/db"
	"main/services/auth/db/pg"
	"main/services/auth/mail"
	"time"
)

//var service Auth

const (
	loginExpirationTime = time.Duration(24 * time.Hour)
)

type Auth struct {
	db   db.Intr
	cash cash.Intr
	mail mail.Intr

	conf Config
}

func (service *Auth) StartService(config Config) (extErr *ExtError.Error) {
	service.conf = config
	service.db = &pg.DBPostgresql{}
	extErr = service.db.Open(config.DbConnect)
	if extErr != nil {
		return extErr
	}

	service.cash = &dbRedis.Db{}
	service.cash.Open(config.CashConfig)

	service.mail = &mail.Mail{}
	extErr = service.mail.Init(config.MailServer, config.MailPort, config.MailPassword, config.MailFrom)

	return extErr
}

func (service *Auth) SignUp(su signUp.Request) (extErr *ExtError.Error) {
	/*if service.db == interfaceDb.Intr(nil) {
		return ExtError.New("Request, Intr not set", 0)
	}*/
	activationKey, extErr := service.db.SignUp(su.Login, su.Email, su.Password)
	if extErr != nil {
		return
	}

	msg := fmt.Sprintf(resource.TmplEmailBoby, su.Login, "http://127.0.0.1:8080/activate?login="+base64.StdEncoding.EncodeToString([]byte(su.Login))+"&data="+activationKey)

	extErr = service.mail.SendEmail(su.Email, "Activation account", msg)
	if extErr != nil {
		service.db.RemoveAccount(su.Email)
		return
	}
	return
}

func (service *Auth) Login(au auth.Request) (status int, extErr *ExtError.Error) {

	if au.SessionId != "" {
		active, extErr := service.cash.GetLogin(au.Login, au.SessionId, au.SessionToken)
		if extErr != nil {
			return 0, extErr
		}
		if active {
			return 0, extErr
		}
	}

	status, extErr = service.db.Login(au.Login, au.Password)
	if extErr == nil {
		service.cash.Login(au.Login, loginExpirationTime)
	}

	return service.db.Login(au.Login, au.Password)
}
