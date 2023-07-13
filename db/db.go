package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	"transactionroutine/configs"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func GenerateSeed() {
	sqlFile := "db/initialData.sql"
	content, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		panic("Failed to read SQL file: " + err.Error())
	}

	queries := strings.TrimSpace(string(content))
	_, err = db.Exec(queries)
	if err != nil {
		panic("Failed to execute SQL statements: " + err.Error())
	}
}

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

	if configs.GetSeedingOpt() == true {
		GenerateSeed()
	}
}

func GetConn() *sql.DB {
	return db
}
