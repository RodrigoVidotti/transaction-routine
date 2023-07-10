package db

import (
	"database/sql"
	"fmt"
	"transactionroutine/configs"
)

var db *sql.DB

func Init() {
	db, err := sql.Open(configs.GetDBType(), configs.GetConnString())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("db connection unavailable")
		panic(err.Error())
	}
	fmt.Println("db connected")
}

func CreateConn() *sql.DB {
	return db
}
