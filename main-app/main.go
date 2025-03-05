package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/revandpratama/reflect/adapter"
	"github.com/revandpratama/reflect/cmd"
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
	minioOption := &adapter.MinioOption{}
	postgresOption := &adapter.PostgresOption{}
	restOption := &adapter.RestOption{}
	grpcOption := &adapter.GRPCOption{}
	a, err := adapter.NewAdapter(
		kafkaOption,
		minioOption,
		postgresOption,
		grpcOption,
		restOption,
	)
	if err != nil {
		server.errorOccured <- err
	}

	if config.ENV.AppEnvironment == "development" {
		helper.NewLog().Info("server running in development mode").ToKafka()

		cmd.AutoMigrate(a.Postgres)
	} else {
		helper.NewLog().Info(fmt.Sprintf("server running in %v mode", config.ENV.AppEnvironment)).ToKafka()
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
		restOption,
		grpcOption,
		postgresOption,
		minioOption,
		kafkaOption,
	)

	helper.NewLog().Info("resources cleaned up")
}
