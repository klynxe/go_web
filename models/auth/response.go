package auth

const (
	OK              = 0
	CAPTCHA_FAILURE = 1
	AUTH            = 2
)

type Response struct {
	Error  int    `json:"error"`
	Msg    string `json:"msg"`
	Status int    `json:"status"`
}
