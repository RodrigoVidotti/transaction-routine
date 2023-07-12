package db

import (
	"database/sql"
	"fmt"
	"time"
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

	for db.Ping() != nil {
		fmt.Println("db connection unavailable")
		time.Sleep(2 * time.Second)
	}
	fmt.Println("db connected")
}

func GetConn() *sql.DB {
	return db
}
