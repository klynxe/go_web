package dbRedis

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/go-redis/redis"
	"main/ExtError"
	"main/rand"
	"time"
)

type Db struct {
	client *redis.Client
}

func (db *Db) Open(url string) (extErr *ExtError.Error) {
	options, err := redis.ParseURL(url)
	if err != nil {
		return ExtError.Resend("Error open redis", 1, err)
	}
	db.client = redis.NewClient(options)

	pong, err := db.client.Ping().Result()
	if err != nil {
		return ExtError.Resend("Error redis ping", 1, err)
	}
	fmt.Println(pong, err)

	err = db.client.Set("sessions:login:session_id", "token", 24*time.Hour).Err()
	if err != nil {
		return ExtError.Resend("Error redis save session", 1, err)
	}
	return
}

func (db *Db) IsOpen() (open bool) {
	if db.client == nil {
		return false
	}
	if err := db.client.Ping().Err(); err != nil {
		return false
	}
	return true
}

func (db *Db) Login(login string, expiration time.Duration) (sessionId, token string, extErr *ExtError.Error) {
	token = rand.String(128)
	md := md5.Sum([]byte(token))
	hexToken := hex.EncodeToString(md[:])
	err := db.client.Set("sessions:"+login+":"+sessionId, hexToken, expiration).Err()
	if err != nil {
		panic("not ready")
		//return ExtError.Resend("Error redis save session", 1, err)
	}
	return
}

func (db *Db) GetLogin(login, sessionId, token string) (sessionActive bool, extErr *ExtError.Error) {
	dbToken, err := db.client.Get("sessions:" + login + ":" + sessionId).Result()
	if err != nil {
		extErr = ExtError.Resend("Error redis save session", 1, err)
		return
	}

	md := md5.Sum([]byte(token))
	hexToken := hex.EncodeToString(md[:])

	if dbToken != hexToken {
		extErr = ExtError.Resend("Error check token", 1, err)
		return
	}

	return
}
