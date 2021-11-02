// BWOTSHEWCHb

package main

import (
	"fmt"

	"github.com/Manni-MinM/Leprechaun/db"
	"github.com/Manni-MinM/Leprechaun/model"
)

func main() {
	db.New()
	db.CreateTable()

	link1 := model.GetLink("google.com")
	link2 := model.GetLink("aut.ac.ir")
/*
	db.InsertRecord(link1)
	db.InsertRecord(link2)
*/
	var err error
	var linkRes model.Link
	linkRes , err = db.SelectRecord("eiJvTO4a")
	if err == nil {
		fmt.Println(linkRes)
	} else {
		fmt.Println(err)
	}

	linkRes , err= db.SelectRecord("HVkg9LRL")
	if err == nil {
		fmt.Println(linkRes)
	} else {
		fmt.Println(err)
	}

	_ = link1
	_ = link2
	_ = err
	_ = linkRes
}

