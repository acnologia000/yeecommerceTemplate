package main

import (
	"log"
	"net/http"
)

// just set it to false for production builds
const DEBUG = true

func main() {

	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	
	// increasing connection pool size for less recycle overhead
	db.SetMaxOpenConns(12)

	http.HandleFunc("/login", login(db))
	http.HandleFunc("/register", register(db))
	http.ListenAndServe(LISTEN_ADDRESS, nil)
}
