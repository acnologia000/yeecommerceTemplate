package main

import (
	"testing"
	"time"
)

const (
	testmail_1 = "some@gmail.com"
	testmail_3 = "some@imail.com"
	testmail_2 = "some@hmail.com"
)

func TestAddAndDeletesession(t *testing.T) {
	key, err := addSession(testmail_1)

	if err != nil {
		t.Error(err)
	}

	t.Logf("testing for : %s", key)

	mail, present := getSession(key)

	if !present {
		t.Error("key not found")
	}

	if mail != testmail_1 {
		t.Error("wrong email")
	}

}

func TestExpireSession(t *testing.T) {

	go expireSession(time.NewTicker(SESSION_CYCLE_TIME))

	key, err := addSession(testmail_1)
	if err != nil {
		t.Error(err)
	}

	t.Log(key)

	time.Sleep(time.Second * 30)

	t.Log("exiting")

	_, Exists := getSession(key)

	if Exists {
		t.Error("sessions not cleaned up")
	}

	t.Log(SessionStore)
}
