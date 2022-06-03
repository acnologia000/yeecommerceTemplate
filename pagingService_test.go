package main

import (
	"fmt"
	"testing"
)

func TestInitPaging(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Fatal(err)
	}
	if err := initPaging(db); err != nil {
		t.Fatal(err)
	}
}

func TestChunkArrayIntoPages(t *testing.T) {
	samplePool := make([]items, 0)

	for i := 0; i < PAGE_ITEM_COUNT*5; i++ {
		samplePool = append(samplePool, items{Name: "", Price: i, Description: ""})
	}

	t.Log(len(samplePool))
	// for 5 pages
	pages := chunkArrayIntoPages(samplePool)
	t.Log(len(pages))
	fmt.Print(pages[1])

	// for perfect 4 pages
	pages = chunkArrayIntoPages(samplePool[0:79])
	if len(pages) != 4 {
		t.Error("fail for 0:79\n")
		t.Fail()
	}
	t.Log("pass for perfect 4\n")
	// for imperfect 4
	pages = chunkArrayIntoPages(samplePool[0:76])
	if len(pages) != 4 {
		t.Error("fail for 0:77")
		t.Fail()
	}

	t.Log("pass for imperfect 4\n")
	// for imperfect 3
	pages = chunkArrayIntoPages(samplePool[0:55])
	if len(pages) != 3 {
		t.Error("fail for 0:55")
		t.Fail()
	}

	t.Log(pages)

	t.Log("pass for imperfect 3\n")
	// for smaller than page size
	pages = chunkArrayIntoPages(samplePool[0 : PAGE_ITEM_COUNT-4])
	if len(pages) != 1 {
		t.Error("fail for smaller than page size")
		t.Fail()
	}

	t.Log("pass for smaller than page size\n")
}
