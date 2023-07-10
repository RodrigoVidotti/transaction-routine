package configs

import "fmt"

const (
	DBHost = "localhost"
	DBPort = "3306"
	DBUser = "root"
	DBPass = "123abc"
	DBName = "mysql"
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
