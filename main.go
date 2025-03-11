package main

import (
	"goserver/pkg/mysql"
	"goserver/pkg/server"
)

func main() {
	mysql.Connect()
	server.Run()
}
