package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math"
	"sync"
	"time"
)

var Pages [][]byte

var plock sync.Mutex

type items struct {
	Name        string `json:"name"`
	Id          string `json:"id"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	ImageID     string `json:"imageID"`
}

func initPaging(db *sql.DB) error {

	if err := populatePages(db); err != nil {
		return err
	}

	ticker := time.NewTicker(PAGE_LIFESPAN)
	InternelRefreshService(*ticker, db)
	return nil
}

func populatePages(db *sql.DB) error {
	rows, err := db.Query("select * from items order by priority desc")

	if err != nil {
		return err
	}
	var entries = make([]items, 1)

	for rows.Next() {
		var name string
		var id string
		var price int
		var description string
		var imageID string
		err = rows.Scan(&name, &id, &price, &description, &imageID)
		if err != nil {
			return err
		}

		entries = append(entries, items{Name: name, Id: id, Price: price, Description: description, ImageID: imageID})
	}

	pages := chunkArrayIntoPages(entries)
	pdata := make([][]byte, 0)

	for i := 0; i < len(pages); i++ {
		t, err := json.Marshal(pages[i])
		if err != nil {
			return err
		}
		pdata = append(pdata, t)
	}
	plock.Lock()
	Pages = pdata
	plock.Unlock()
	return nil
}

func InternelRefreshService(t time.Ticker, db *sql.DB) {
	for range t.C {
		populatePages(db)
	}
}

func chunkArrayIntoPages(it []items) [][]items {
	var (
		totalItems = len(it)
		returnVal  = make([][]items, 0)
	)
	// in case nummber of items in data is smaller than page item count
	// we can just pack the input in another slice and return without
	// actually wasting CPU or memory
	if totalItems < PAGE_ITEM_COUNT {
		print("total items less than page count")
		returnVal = append(returnVal, it)
		return returnVal
	}

	var (
		pageCount = int(math.Floor(float64(totalItems) / float64(PAGE_ITEM_COUNT)))
		remaining = totalItems - (pageCount * PAGE_ITEM_COUNT)
		prev      = 0
		next      = PAGE_ITEM_COUNT
	)

	for i := 0; i < pageCount; i++ {
		temp := it[prev:next]
		returnVal = append(returnVal, temp)
		prev = next
		next = next + PAGE_ITEM_COUNT
	}
	fmt.Printf("lenght of return value is %d ", len(returnVal))

	fmt.Printf("\nremaining = %d\n", remaining)
	if remaining != 0 {
		lastIndex := totalItems - 1
		print("remaining fired")
		// rewinding a step back to get to the last page just afterlast counted
		// item hence (PAGE_ITEM_COUNT-1) instead of PAGE_ITEM_COUNT
		lastCountedIndex := (totalItems - remaining)
		returnVal = append(returnVal, it[lastCountedIndex:lastIndex])
	}

	return returnVal
}
