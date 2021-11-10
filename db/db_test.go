// BWOTSHEWCHb

package db

import (
	"testing"

	"github.com/Manni-MinM/Leprechaun/model"
)

func TestDB(t *testing.T) {
	var err error
	err = New()
	if err != nil {
		t.Error(err)
	}
	err = CreateTable()
	if err != nil {
		t.Error(err)
	}
	
	var links []model.Link
	google := model.GetLink("google.com")
	links = append(links , google)
	
	aut := model.GetLink("aut.ac.ir")
	aut.SetHash("Bloody Hell")
	links = append(links , aut)

	github := model.GetLink("github.com")
	links = append(links , github)

	for _ , link := range(links) {
		err = InsertRecord(link , "one_day") 
		if err != nil {
			t.Error("Insertion Error")
		}
	}

	for _ , link := range(links) {
		row , _ := SelectRecord(link.Hash)
		if row != link {
			t.Error("Validation Error")
		}
	}

	for _ , link := range(links) {
		rowBefore , _ := SelectRecord(link.Hash)
		err = UpdateRecord(link)
		rowAfter , _ := SelectRecord(link.Hash)
		if rowAfter.UsedCount - rowBefore.UsedCount != 1 {
			t.Error("Update Error")
		}
	}

	for _ , link := range(links) {
		err = ExpireRecord(link.Hash)
		_ , errSel := SelectRecord(link.Hash) 
		if errSel != nil {
			t.Error("Expiration Error")
		}
	}

	err = DropTable()
	if err != nil {
		t.Error(err)
	}
}

