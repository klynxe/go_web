package resource

const (
	TmplEmail = "From:<%s>\r\n" +
		"To: <%s>\r\n" +
		"Subject: %s\r\n" +
		"\r\n" +
		"%s"

	TmplEmailBoby = "Hello dear %s,\r\n" +
		"to activation your account on our site please follow the link\r\n" +
		"%s"
)
