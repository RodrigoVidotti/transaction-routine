package db

import (
	"database/sql"
	"fmt"
	"transactionroutine/configs"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init() {
	var err error
	db, err = sql.Open(configs.GetDBType(), configs.GetConnString())
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("db connection unavailable")
		panic(err.Error())
	}
	fmt.Println("db connected")
}

func GetConn() *sql.DB {
	return db
}
