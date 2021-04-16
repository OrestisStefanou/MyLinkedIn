package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func dbConnect() {
	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", databaseUser, userPassword, databaseName))
	checkErr(err)

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
