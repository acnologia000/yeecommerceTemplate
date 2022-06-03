package main

import (
	"encoding/json"
	"strings"
	"testing"
)

const sampleDateFloat = 20110101 // YYYYMMDD
const sampleSearchJson = `
{
	"MaxPrice":333,
	"MinPrice":555,
	"PriceOrder":1,
	"ldate":20130101,
	"edate":20200101,
	"PriceOrder":-1
}
`

func TestPrepateSearchQuery(t *testing.T) {
	data := make(map[string]float64, 0)
	err := json.Unmarshal([]byte(sampleSearchJson), &data)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(PrepareSearchQuery(data))
}

func TestOrderByParameterSQL(t *testing.T) {
	var base strings.Builder
	var pref = make(map[string]string)
	_, err := base.WriteString("   ")

	if err != nil {
		t.Fatal(err)
	}
	// for empty map
	orderByParameterSQL(pref, &base)

	if base.String() != "   " {
		t.Error("-- falied --")
	}

	// for 2 keys
	pref["price"] = "ASC"
	pref["added_on"] = "ASC"

	orderByParameterSQL(pref, &base)

	if res := base.String(); res != "   order by price ASC, added_on ASC" {
		t.Errorf("-- falied -- \n result : \n %s", res)
	}
}

func TestGetSingleItemById(t *testing.T) {
	setupTestEnv()
	connection, err := connect()
	if err != nil {
		t.Fatal(err)
	}
	jsonData, err := getSingleItemById(connection, "")

	if err != nil {
		t.Fatal(err)
	}

	var x items

	err = json.Unmarshal(jsonData, &x)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(x)
}
func TestFloatToDate(t *testing.T) { // passed
	date := IntToDate(sampleDateFloat)
	if date == "1971-01-01" {
		t.Fatal(date)
	}
	t.Log(date)
	print(date)
}
