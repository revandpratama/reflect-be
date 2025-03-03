package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/revandpratama/reflect/adapter"
	"github.com/revandpratama/reflect/config"
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

	err := godotenv.Load()
	if err != nil {
		server.errorOccured <- err
	}

	// * load global config
	config.LoadConfig()

	kafkaOption := &adapter.KafkaGoOption{}
	postgresOption := &adapter.PostgresOption{}
	restOption := &adapter.RestOption{}
	grpcOption:= &adapter.GRPCOption{}
	a, err := adapter.NewAdapter(
		kafkaOption,
		postgresOption,
		restOption,
		grpcOption,
	)
	if err != nil {
		server.errorOccured <- err
	}


	// ? block server until shutdown signals
	select {
	case sig := <-server.shutdown:
		helper.NewLog().Info(fmt.Sprintf("Server shutting down, cause: %v", sig)).ToKafka()
	case err := <-server.errorOccured:
		helper.NewLog().Fatal(fmt.Sprintf("Error starting server, cause: %v", err)).ToKafka()
	}

	helper.NewLog().Info("shutting down server...")
	helper.NewLog().Info("cleaning up resources...")

	a.Close(
		postgresOption,
		restOption,
		grpcOption,
		kafkaOption,
	)

	helper.NewLog().Info("resources cleaned up")
	helper.NewLog().Info("server stopped").ToKafka()
}
