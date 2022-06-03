package main

import (
	"testing"
)

func TestConnect(t *testing.T) {
	setupTestEnv()
	db, err := connect()

	if err != nil {
		t.Error("connection failed")
	}

	if err := db.Ping(); err != nil {
		t.Log(err)
		t.Error("ping failed")
	}
	db.Close()
	// os.Unsetenv("DB_USERNAME")
	// os.Unsetenv("DB_PASSWORD")
	// os.Unsetenv("DB_NAME")
}
