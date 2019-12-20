package auth

import (
	"gopkg.in/go-playground/validator.v9"
	"main/ExtError"
)

type Request struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required_without=SessionId"`

	SessionId    string `json:"session-id" validate:"required_without=Password"`
	SessionToken string `json:"session-token" validate:"required_with=SessionId"`

	CaptchaToken string `json:"captcha-token" validate:"required"`
}

func (re *Request) Validate() (extErr *ExtError.Error) {
	validate := validator.New()
	if err := validate.Struct(re); err != nil {
		extErr = ExtError.New("Error validate auth", 0)
	}
	return
}
