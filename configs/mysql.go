package configs

import "fmt"

const (
	DBHost          = "db"
	DBPort          = "3306"
	DBUser          = "root"
	DBPass          = "password"
	DBName          = "transaction_routine"
	DBType          = "mysql"
	Seeding         = true
	MultiStatements = "true"
)

func GetDBType() string {
	return DBType
}

func GetSeedingOpt() bool {
	return Seeding
}

func GetConnString() string {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=%s",
		DBUser,
		DBPass,
		DBHost,
		DBPort,
		DBName,
		MultiStatements)

	return conn
}
