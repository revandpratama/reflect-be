package adapter

import (
	"fmt"

	"github.com/revandpratama/reflect/config"
	"github.com/revandpratama/reflect/helper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCOption struct {
	GrcpClient *grpc.ClientConn
}

func (g *GRPCOption) Start(a *Adapter) error {

	helper.NewLog().Info("initializing grpc server...").ToKafka()

	conn, err := grpc.NewClient(fmt.Sprintf(":%v", config.ENV.GRPCServerPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to server: %v", err)
	}

	g.GrcpClient = conn
	a.GrcpClient = conn

	helper.NewLog().Info("grpc server started").ToKafka()

	return nil
}

func (g *GRPCOption) Stop() error {

	g.GrcpClient.Close()

	return nil
}
