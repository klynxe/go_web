package cash

import (
	"main/ExtError"
	"time"
)

type Intr interface {
	IsOpen() (open bool)
	Open(connect string) (extErr *ExtError.Error)

	Login(login string, expiration time.Duration) (sessionId, token string, extErr *ExtError.Error)
	GetLogin(login, sessionId, token string) (sessionActive bool, extErr *ExtError.Error)
}
