// BWOTSHEWCHb

package db

import (
	"fmt"
	"errors"
	"database/sql"

	"github.com/Manni-MinM/Leprechaun/model"

	_"github.com/go-sql-driver/mysql"
)

// connects to mysql and creates new database instance
func New() error {
	// initialize variables
	dbDriver := "mysql"
	dbUser := "root"
	dbPassword := "pashmak64bit"
	// connect to mysql
	key := fmt.Sprintf("%s:%s@/" , dbUser , dbPassword)
	db , err := sql.Open(dbDriver , key)
	if err != nil {
		panic(err)
	}
	// create new database
	queryCreate :=
		`CREATE DATABASE IF NOT EXISTS Leprechaun ;`
	rows , err := db.Query(queryCreate)
	_ = rows
	return err
}
// connects to mysql and returns db instance
func Connect() (db *sql.DB) {
	// initialize variables	
	dbDriver := "mysql"
	dbUser := "root"
	dbPassword := "pashmak64bit"
	dbName := "Leprechaun"
	// connect to mysql
	key := fmt.Sprintf("%s:%s@/%s" , dbUser , dbPassword , dbName)
	db , err := sql.Open(dbDriver , key)
	if err != nil {
		panic(err)
	}
	return db
}
// creates new table in db
func CreateTable() error {
	// get db instance
	db := Connect()
	// create new table in db
	query :=
		`CREATE TABLE IF NOT EXISTS Links(
			Hash varchar(255) PRIMARY KEY ,
			URL varchar(255) ,
			UsedCount int DEFAULT 0 ,
			ExpiryDate datetime DEFAULT NOW()
		) ;`
	rows , err := db.Query(query)
	_ = rows
	// close connection to mysql
	db.Close()
	return err
}
// deletes table in db
func DropTable() error {
	// get db instance
	db := Connect()
	// deletes Links table in db
	query :=
		`DROP TABLE Links ;`
	rows , err := db.Query(query)
	// close connection to mysql
	_ = rows
	db.Close()
	return err
}
// deletes expired records from db
func ExpireRecord(hash string) error {
	// get db instance
	db := Connect()
	// deletes record from table if expired
	query := 
		`DELETE FROM Links WHERE (Hash = ? AND ExpiryDate <= NOW()) ;`
	rows , err := db.Query(query , hash)
	_ = rows
	// close connection to mysql
	db.Close()
	return err
}
// insert record into db
func InsertRecord(link model.Link , expiryDate string) error {
	// get db instance
	db := Connect()
	// check if record has expired and return if so
	err := ExpireRecord(link.Hash)
	if err != nil {
		return err
	}
	// insert record into table
	query := ""
	if expiryDate == "one_month" {
		query = 
			`INSERT IGNORE INTO Links(Hash , URL , ExpiryDate) VALUES (? , ? , DATE_ADD(NOW() , INTERVAL 1 MONTH)) ;`
	} else if expiryDate == "one_week" {
		query = 
			`INSERT IGNORE INTO Links(Hash , URL , ExpiryDate) VALUES (? , ? , DATE_ADD(NOW() , INTERVAL 1 WEEK)) ;`
	} else {
		query = 
			`INSERT IGNORE INTO Links(Hash , URL , ExpiryDate) VALUES (? , ? , DATE_ADD(NOW() , INTERVAL 1 DAY)) ;`
	}
	rows , err := db.Query(query , link.Hash , link.URL)
	_ = rows
	// close connection to mysql
	db.Close()
	return err
}
// increments UsedCount parameter of record in db by one
func UpdateRecord(link model.Link) error {
	// get db instance
	db := Connect()
	// update UsedCount on specified record in table
	query := 
		`UPDATE Links SET UsedCount = ? WHERE Hash = ?`
	rows , err := db.Query(query , link.UsedCount + 1 , link.Hash)
	_ = rows
	// close connection to mysql
	db.Close()
	return err
}
// finds record with specified hash and returns it
func SelectRecord(hash string) (model.Link , error) {
	// get db instance
	db := Connect()
	link := model.Link{}
	// check if record has expired and return if so
	err := ExpireRecord(hash)
	if err != nil {
		return link , err
	}
	// select records from table with specified hash
	query := 
		`SELECT Hash , URL , UsedCount FROM Links WHERE Hash = ? ;`
	rows , err := db.Query(query , hash)
	if err != nil {
		panic(err)
	}	
	// close connection to mysql
	db.Close()
	// return link and nil if record exists otherwise return link and error
	for rows.Next() {
		err = rows.Scan(&link.Hash , &link.URL , &link.UsedCount)
		return link , nil
	}
	return link , errors.New("No Such Link or Linked Expired")
}

