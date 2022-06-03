package main

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
)

const loginRequestBody = `
{
	"email":"some@email.com",
	"password":"some password string"
}`

func TestLogin(t *testing.T) {

	setupTestEnv()
	db, err := connect()

	if err != nil {
		t.Error(err)
	}

	// registerig fake user to test login
	reg_req := httptest.NewRequest("POST", "/login", strings.NewReader(body))
	reg_res := httptest.NewRecorder()

	register(db)(reg_res, reg_req)

	var reg_response = make(map[string]string)
	err = json.NewDecoder(reg_res.Body).Decode(&reg_response)

	if err != nil {
		t.Error(err)
	}

	// actual login testing logic begins here
	req := httptest.NewRequest("POST", "/login", strings.NewReader(loginRequestBody))
	res := httptest.NewRecorder()

	login(db)(res, req)

	var response = make(map[string]string)
	err = json.NewDecoder(res.Body).Decode(&response)

	if err != nil {
		t.Error(err)
	}

	cause, failure := response["failure"]

	if failure {
		t.Error(cause)
	}

	t.Log(response["token"])

}
