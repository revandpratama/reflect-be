package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/revandpratama/reflect/helper"
)

type Server struct {
	shutdown     chan os.Signal
	errorOccured chan error
}

func NewServer() *Server {
	return &Server{
		shutdown:     make(chan os.Signal, 1),
		errorOccured: make(chan error, 1),
	}
}

func main() {
	server := NewServer()
	server.start()
}

func (server *Server) start() {
	signal.Notify(server.shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-server.shutdown:
		helper.NewLog().Info(fmt.Sprintf("Server shutting down, cause: %v", sig)).ToKafka()
	case err := <-server.errorOccured:
		helper.NewLog().Fatal(fmt.Sprintf("Error starting server, cause: %v", err)).ToKafka()
	}
}
