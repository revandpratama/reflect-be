package main

import "github.com/revandpratama/reflect/logging-service/server"

func main() {
	server := server.NewServer()
	server.Start()
}