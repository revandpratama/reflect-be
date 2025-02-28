package adapter

import (
	"fmt"
	"net"

	"github.com/revandpratama/reflect/auth-service/helper"
	"google.golang.org/grpc"
)

type GRPCOption struct {
	GrcpServer *grpc.Server
}

func (g *GRPCOption) Start(a *Adapter) error {

	helper.NewLog().Info("initializing grpc server...").ToKafka()

	listener, err := net.Listen("tcp", ":50051") // Adjust port as needed
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	g.GrcpServer = grpc.NewServer()

	go func() {

		if err := g.GrcpServer.Serve(listener); err != nil {
			// Handle server failure
			fmt.Printf("Failed to serve gRPC: %v\n", err)
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
