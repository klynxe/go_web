package auth

type Config struct {
	DbConnect string

	CashConfig string

	MailServer   string
	MailPort     string
	MailPassword string
	MailFrom     string
}
