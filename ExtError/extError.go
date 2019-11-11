package ExtError

type Error struct {
	msg     string
	code    int
	errPrev error
}

func (p *Error) Error() string {
	if p.errPrev != nil {
		return p.msg + " (errPrev = " + p.errPrev.Error() + ")"
	}
	return p.msg
}

func New(msg string, code int) *Error {
	return &Error{msg: msg, code: code}
}

func Resend(msg string, code int, errPrev error) *Error {
	return &Error{msg: msg, code: code, errPrev: errPrev}
}
