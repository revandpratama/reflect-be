package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/revandpratama/reflect/auth-service/adapter"
	"github.com/revandpratama/reflect/auth-service/config"
	"github.com/revandpratama/reflect/auth-service/helper"
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

	err := godotenv.Load()
	if err != nil {
		server.errorOccured <- err
	}

	// * load global config
	config.LoadConfig()

	// * initialize adapters
	a, err := adapter.NewAdapter(&adapter.PostgresOption{})
	if err != nil {
		server.errorOccured <- err
	}

	helper.NewLog("INFO", "server running.")
	select {
	case sig := <-server.shutdown:
		helper.NewLog("INFO", fmt.Sprintf("Server shutting down, cause: %v", sig))
	case err := <-server.errorOccured:
		helper.NewLog("FATAL", fmt.Sprintf("Error starting server, cause: %v", err))
	}

	server.cleanup(a)
}

func (s *Server) start() {
	signal.Notify(s.shutdown, os.Interrupt, syscall.SIGTERM)

}

func (s *Server) cleanup(a *adapter.Adapter) {
	helper.NewLog("INFO", "cleaning up resources...")

	a.Close(
		&adapter.PostgresOption{},
	)

	helper.NewLog("INFO", "resource cleaned up")
}
