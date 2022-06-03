package main

import (
	"bytes"
	"net/http/httptest"
	"testing"
)

func TestPage(t *testing.T) {
	conn, err := connect()
	if err != nil {
		t.Fatal(err)
	}
	// for invalid page number
	container := bytes.NewBuffer(make([]byte, 0))
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/page/x", container)
	page(conn)(res, req)
}
