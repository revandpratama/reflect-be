package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
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

	select {
	case sig := <-server.shutdown:
		log.Info().Msg(fmt.Sprintf("Server shutting down, cause: %v", sig))
	case err := <-server.errorOccured:
		log.Error().Msg(fmt.Sprintf("Error starting server, cause: %v", err))
	}

	server.cleanup()
}

func (s *Server) start() {
	signal.Notify(s.shutdown, os.Interrupt, syscall.SIGTERM)

}

func (s *Server) cleanup() {

}
