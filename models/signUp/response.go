package signUp

const (
	OK              = 0
	CAPTCHA_FAILURE = 1
	SIGN_UP         = 2
)

type Response struct {
	Error int    `json:"error"`
	Msg   string `json:"msg"`
}
