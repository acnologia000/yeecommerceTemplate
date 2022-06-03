package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func connect() (*sql.DB, error) {

	username, isPresent := os.LookupEnv("DB_USERNAME")

	print(isPresent)
	printDebug("using username :%s\n", username)

	if !isPresent {
		log.Fatal("username not set, please set DB_USERNAME environment variable")
	}

	password, isPresent := os.LookupEnv("DB_PASSWORD")

	if !isPresent {
		log.Fatal("password not set, please set DB_USERNAME environment variable")
	}

	dbName, isPresent := os.LookupEnv("DB_NAME")

	if !isPresent {
		log.Fatal("database name not set, please set DB_USERNAME environment variable")
	}

	printDebug("using database :%s\n", dbName)

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", username, dbName, password)
	return sql.Open("postgres", connStr)

}
