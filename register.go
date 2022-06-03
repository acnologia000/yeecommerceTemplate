package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func register(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			w.Write(errorResponse(INVALID_METHOD))
			return
		}

		var data = make(map[string]string, 1)

		if data == nil {
			log.Fatal("data is nil")
		}

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			w.Write(errorResponse(JSON_DECODE_ERROR))
			return
		}
		print("decode success")

		defer r.Body.Close()

		stmt, err := db.Prepare(INSERT_NEW_USER)

		if err != nil {
			w.Write(errorResponse(STMT_PREPARE_ERROR))
			return
		}

		result, err := stmt.Exec(data["name"], data["email"], ToBaseEncodedMD5(data["password"]))

		if err != nil {
			log.Print(err)
			w.Write(errorResponse(QUERY_ERROR))
			return
		}

		rowsAffected, err := result.RowsAffected()

		if err != nil {
			log.Print(err)
			w.Write(errorResponse(QUERY_ERROR))
			return
		}

		log.Printf("rows affected:%d", rowsAffected)

		w.Write([]byte(SUCCESS_RESPONSE))
	}
}
