package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func getSingleItemById(db *sql.DB, name string) ([]byte, error) {
	stmt, err := db.Prepare(SEARCH_ITEM_QUERY)

	if err != nil {
		return nil, err
	}

	row, err := stmt.Query(name)

	if err != nil {
		return nil, err
	}

	if !row.Next() {
		return nil, errors.New("no result")
	}

	var iname, description, id, imageID string
	var price int

	err = row.Scan(&iname, &id, &price, &description, &imageID)

	if err != nil {
		return nil, err
	}

	itemData := items{Name: iname, Id: id, Price: price, Description: description, ImageID: imageID}
	Json, err := json.Marshal(itemData)

	if err != nil {
		return nil, err
	}

	return Json, nil
}

/*
order and range by number units only

price
rank
release Date

limit/offset are in multiple of PAGE_ITEM_COUNT,
search results are paged

> if both min and max are 0 then whole parameter is dropped
> for *Order
	--any postive non 0 value means ascending,
	--any 0 value means, to not order by the parameter
	--any negative non 0 value means descending order
*/

// this function has to be reworked
func PrepareSearchQuery(parameters map[string]float64) string {
	var (
		Base    strings.Builder
		OrderBy = make(map[string]string)
		//price parameters processing starts here
		MaxPrice   = parameters["MaxPrice"]
		MinPrice   = parameters["MinPrice"]
		PriceOrder = int(math.Floor(parameters["PriceOrder"]))
	)
	Base.WriteString("select * from items where ")
	if MaxPrice != 0 && MinPrice != 0 {
		if MaxPrice == 0 {
			MaxPrice = math.MaxFloat32
		}
		Base.WriteString(fmt.Sprintf("(price between %f and %f) and ", MaxPrice, MinPrice))
	}

	if PriceOrder > 0 {
		OrderBy["price"] = "ASC"
	} else if PriceOrder < 0 {
		OrderBy["price"] = "DSC"
	}

	// time stamp processing starts here
	var (
		latest    = IntToDate(int(math.Floor(parameters["ldate"])))
		earliest  = IntToDate(int(math.Floor(parameters["edate"])))
		dateOrder = int(parameters["dOrder"])
	)

	if latest != earliest {
		Base.WriteString(fmt.Sprintf("(added_on between TO_TIMESTAMP('%s', 'YYYY-MM-DD') and TO_TIMESTAMP('%s', 'YYYY-MM-DD')) and ", latest, earliest))
	}

	if dateOrder > 0 {
		OrderBy["added_on"] = "ASC"
	} else if PriceOrder < 0 {
		OrderBy["added_on"] = "DSC"
	}

	orderByParameterSQL(OrderBy, &Base)

	return Base.String()
}

func orderByParameterSQL(input map[string]string, base *strings.Builder) {
	keyCount := len(input)
	print(keyCount)
	counter := 1
	if keyCount == 0 {
		return
	}

	base.WriteString("order by")

	for k, v := range input {
		base.WriteString(" ")
		base.WriteString(k)
		base.WriteString(" ")
		base.WriteString(v)

		if counter < keyCount {
			print(keyCount, "  ", counter, "\n")
			base.WriteByte(',')
		}

		counter++
	}

}

func IntToDate(data int) string {
	str := strconv.FormatInt(int64(data), 10)

	if len(str) != 8 {
		return "1971-01-01"
	}
	var (
		year  = str[0:4]
		month = str[4:6]
		day   = str[6:8]
	)
	return fmt.Sprintf("%s-%s-%s", year, month, day)
}
