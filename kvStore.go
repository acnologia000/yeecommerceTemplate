package main

import (
	"log"
	"sync"
	"time"
)

type LoginData struct {
	email     string
	loginTime time.Time
}

// key = secret code / token
var SessionStore = make(map[string]LoginData)

// session data mutex

var SessionMutex = &sync.RWMutex{}

func addSession(email string) (string, error) {
	SessionMutex.Lock()
	key, err := GenerateRandomString(265)
	SessionStore[key] = LoginData{email: email, loginTime: time.Now()}
	SessionMutex.Unlock()
	return key, err
}

func getSession(token string) (string, bool) {
	SessionMutex.Lock()
	data, isPresent := SessionStore[token]
	SessionMutex.Unlock()
	return data.email, isPresent
}

func expireSession(tik *time.Ticker) {
	for t := range tik.C {
		log.Printf("expiring sessions at %s", t.Local().Format(time.RFC1123))

		SessionMutex.Lock()
		for k, v := range SessionStore {
			print(time.Since(v.loginTime))
			if SESSION_DURATION > time.Since(v.loginTime) {
				printDebug("deleting session")
				delete(SessionStore, k)
			}
		}
		SessionMutex.Unlock()
	}

	print("expire session service exit")
}
