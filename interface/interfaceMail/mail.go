package interfaceMail

import "main/ExtError"

type SMTP interface {
	Init(host, port, password, from string) (extErr *ExtError.Error)
	SendEmail(email, msg string) (extErr *ExtError.Error)
}
