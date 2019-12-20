package mail

import "main/ExtError"

type Intr interface {
	Init(host, port, password, from string) (extErr *ExtError.Error)
	SendEmail(email, them, body string) (extErr *ExtError.Error)
}
