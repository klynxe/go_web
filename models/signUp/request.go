package signUp

import (
	"gopkg.in/go-playground/validator.v9"
	"main/ExtError"
)

type Request struct {
	Login    string `json:"login" validate:"required,min=5,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=128"`
	Token    string `json:"token" validate:"required"`
}

func (re *Request) Validate() (extErr *ExtError.Error) {
	validate := validator.New()
	if err := validate.Struct(re); err != nil {
		extErr = ExtError.New("Error validate sign up", 0)
	}
	return
}
