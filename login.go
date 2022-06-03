package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func login(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		//----------------- reading and sanitizing request ----------------------------------
		data := make(map[string]string)

		if r.Method != http.MethodPost {
			w.Write(errorResponse(ILLEGAL_REQUEST_METHOD))
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			w.Write(errorResponse(JSON_DECODE_ERROR))
			return
		}

		defer r.Body.Close()

		//-------------------- extracting data from sql server (storing) -------------------------------
		stmt, err := db.Prepare(LOGIN_SELECT)

		if err != nil {
			w.Write(errorResponse(STMT_PREPARE_ERROR))
			return
		}

		row, err := stmt.Query(data["email"])

		if err != nil {
			log.Print(err)
			w.Write(errorResponse(QUERY_ERROR))
			return
		}
		var hash string

		if !row.Next() {
			w.Write(errorResponse("creds not found"))
			return
		}

		err = row.Scan(&hash)

		if err != nil {
			log.Print(err)
			log.Print("scan error")
			w.Write(errorResponse(QUERY_ERROR))
			return
		}

		//-------------------comparing and deciding response -----------------------------------------
		if ToBaseEncodedMD5(data["password"]) != hash {
			w.Write(errorResponse(PASSWORD_ERROR))
			return
		}

		token, err := addSession(data["email"])

		if err != nil {
			w.Write(errorResponse(INTERNAL_SERVER_ERROR))
			return
		}

		FinalResponse, err := json.Marshal(map[string]string{"token": token})

		if err != nil {
			w.Write(errorResponse(INTERNAL_SERVER_ERROR))
			return
		}

		w.Write(FinalResponse)

	}
}
