package configs

import "fmt"

const (
	DBHost = "db"
	DBPort = "3306"
	DBUser = "root"
	DBPass = "password"
	DBName = "transaction_routine"
	DBType = "mysql"
)

func GetDBType() string {
	return DBType
}

func GetConnString() string {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		DBUser,
		DBPass,
		DBHost,
		DBPort,
		DBName)

	return conn
}
