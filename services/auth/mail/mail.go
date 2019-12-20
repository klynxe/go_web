package mail

import (
	"fmt"
	"main/ExtError"
	"main/resource"
	"net/smtp"
)

type Mail struct {
	host     string
	port     string
	password string
	from     string
	auth     smtp.Auth
}

func (m *Mail) Init(host, port, password, from string) (extErr *ExtError.Error) {
	m.host = host
	m.port = port
	m.password = password
	m.from = from
	m.auth = smtp.PlainAuth("", from, password, host)
	return
}

func (m *Mail) SendEmail(email, them, body string) (extErr *ExtError.Error) {

	msg := fmt.Sprintf(resource.TmplEmail, m.from, email, them, body)

	if err := smtp.SendMail(m.host+":"+m.port, m.auth, m.from, []string{email}, []byte(msg)); err != nil {
		return ExtError.Resend("Error send email", 1, err)
	}
	return
}
