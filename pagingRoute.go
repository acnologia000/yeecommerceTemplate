package main

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"
)

func page(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Write(page_get(db, parsePageNumber(r.URL.Path)))
		case http.MethodDelete:
			w.Write(page_del(db))
		default:
			w.Write(errorResponse(INVALID_METHOD))
		}
	}
}

func page_get(db *sql.DB, pageNumber int) []byte {
	if pageNumber > len(Pages) || pageNumber < 0 {
		return errorResponse(INVALID_PAGE_NUMBER)
	}
	return Pages[pageNumber]
}

func page_del(db *sql.DB) []byte {
	populatePages(db)
	return []byte(`DONE`)
}

func parsePageNumber(url string) int {
	slc := strings.Split(url, "/")
	if len(slc) < 1 {
		return -1
	}

	i, err := strconv.Atoi(slc[len(slc)-1])

	if err != nil {
		return -1
	}
	return i
}
