package main

import (
	"transactionroutine/db"
	"transactionroutine/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start("localhost:8080"))
}
