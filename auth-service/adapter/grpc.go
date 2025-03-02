package adapter

import (
	"fmt"
	"net"
	"os"

	"github.com/revandpratama/reflect/auth-service/config"
	"github.com/revandpratama/reflect/auth-service/helper"
	"github.com/revandpratama/reflect/auth-service/internal/controller"
	pb "github.com/revandpratama/reflect/auth-service/internal/generatedProtobuf/auth"
	"github.com/revandpratama/reflect/auth-service/internal/repositories"
	"github.com/revandpratama/reflect/auth-service/internal/services"
	"google.golang.org/grpc"
)

type GRPCOption struct {
	GrcpServer *grpc.Server
}

func (g *GRPCOption) Start(a *Adapter) error {

	helper.NewLog().Info("initializing grpc server...").ToKafka()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", config.ENV.GRPCServerPort)) // Adjust port as needed
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	g.GrcpServer = grpc.NewServer()

	repo := repositories.NewAuthRepository(a.Postgres)
	service := services.NewAuthService(repo)

	pb.RegisterAuthServiceServer(g.GrcpServer, controller.NewAuthController(service))

	go func() {

		if err := g.GrcpServer.Serve(listener); err != nil {
			// Handle server failure
			helper.NewLog().Fatal(fmt.Sprintf("Failed to start gRPC server: %v", err)).ToKafka()
			os.Exit(1)
		}

	}()

	a.GrcpServer = g.GrcpServer

	helper.NewLog().Info("grpc server started").ToKafka()

	return nil
}

func (g *GRPCOption) Stop() error {

	g.GrcpServer.GracefulStop()

	return nil
}
