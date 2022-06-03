package main

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
)

const body = `
{	
	"name":"random name",
	"email":"some@email.com",
	"password":"some password string"
}`

func TestRegister(t *testing.T) {
	setupTestEnv()

	db, err := connect()

	if err != nil {
		t.Error(err)
	}

	req := httptest.NewRequest("POST", "/login", strings.NewReader(body))
	res := httptest.NewRecorder()

	register(db)(res, req)

	var response = make(map[string]string)
	err = json.NewDecoder(res.Body).Decode(&response)

	if err != nil {
		t.Error(err)
	}

	cause, failure := response["failure"]

	if failure {
		t.Error(cause)
	}

	_, err = db.Exec("delete from users")

	if err != nil {
		t.Log("cleanup failed")
		t.Error(err)
	}

}
