package main

import (
	"net"
	"os"
	"transactionroutine/db"
	"transactionroutine/routes"
)

func getContainerIP() string {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	addrs, err := net.LookupHost(hostname)
	if err != nil {
		panic(err)
	}

	return addrs[0]
}

func main() {
	db.Init()
	e := routes.Init()

	ip := getContainerIP()

	e.Logger.Fatal(e.Start(ip + ":8080"))
}
