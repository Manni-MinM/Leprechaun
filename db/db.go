// BWOTSHEWCHb

package db

import (
	"fmt"
	"errors"
	"database/sql"

	"github.com/Manni-MinM/Leprechaun/model"

	_"github.com/go-sql-driver/mysql"
)

func New() error {
	dbDriver := "mysql"
	dbUser := "root"
	dbPassword := "pashmak64bit"

	key := fmt.Sprintf("%s:%s@/" , dbUser , dbPassword)
	db , err := sql.Open(dbDriver , key)
	if err != nil {
		panic(err)
	}

	queryCreate :=
		`CREATE DATABASE IF NOT EXISTS Leprechaun ;`
	rows , err := db.Query(queryCreate)
	_ = rows
	return err
}
func Connect() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPassword := "pashmak64bit"
	dbName := "Leprechaun"

	key := fmt.Sprintf("%s:%s@/%s" , dbUser , dbPassword , dbName)
	db , err := sql.Open(dbDriver , key)
	if err != nil {
		panic(err)
	}
	return db
}
func CreateTable() error {
	db := Connect()
	query :=
		`CREATE TABLE IF NOT EXISTS Links(
			Hash varchar(255) PRIMARY KEY ,
			URL varchar(255) ,
			UsedCount int DEFAULT 0 ,
			ExpiryDate datetime DEFAULT NOW()
		) ;`
	rows , err := db.Query(query)
	_ = rows
	db.Close()
	return err
}
func DropTable() error {
	db := Connect()
	query :=
		`DROP TABLE Links ;`
	rows , err := db.Query(query)
	_ = rows
	db.Close()
	return err
}
func ExpireRecord(hash string) error {
	db := Connect()
	query := 
		`DELETE FROM Links WHERE (Hash = ? AND ExpiryDate <= NOW()) ;`
	rows , err := db.Query(query , hash)
	_ = rows
	db.Close()
	return err
}
func InsertRecord(link model.Link) error {
	db := Connect()
	err := ExpireRecord(link.Hash)
	if err != nil {
		return err
	}
	query := 
		`INSERT INTO Links(Hash , URL , ExpiryDate) VALUES (? , ? , DATE_ADD(NOW() , INTERVAL 1 MINUTE)) ;`
	rows , err := db.Query(query , link.Hash , link.URL)
	_ = rows
	db.Close()
	return err
}
func UpdateRecord(link model.Link) error {
	db := Connect()
	query := 
		`UPDATE Links SET UsedCount = ? WHERE Hash = ?`
	rows , err := db.Query(query , link.UsedCount + 1 , link.Hash)
	_ = rows
	db.Close()
	return err
}
func SelectRecord(hash string) (model.Link , error) {
	db := Connect()
	link := model.Link{}
	err := ExpireRecord(hash)
	if err != nil {
		return link , err
	}
	query := 
		`SELECT Hash , URL , UsedCount FROM Links WHERE Hash = ? ;`
	rows , err := db.Query(query , hash)
	if err != nil {
		panic(err)
	}
	db.Close()
	if rows.Next() {
		err = rows.Scan(&link.Hash , &link.URL , &link.UsedCount)
		return link , nil
	} else {
		return link , errors.New("No Such Link or Linked Expired")
	}
}

