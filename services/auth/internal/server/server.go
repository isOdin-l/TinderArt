package server

import (
	"context"
	"time"

	config "github.com/isOdin-l/TinderArt/services/auth/configs"
	grpc_gen "github.com/isOdin-l/TinderArt/services/auth/pkg/grpc/grpc-gen"
	"github.com/labstack/echo/v5"
	"google.golang.org/grpc"
)

// Creation of gRPC client to connect with Profile service
type GrpcClient struct {
	Conn   *grpc.ClientConn
	Client grpc_gen.AuthServiceClient
}

func NewGrpcClient(cfg *config.ServerConfig) (*GrpcClient, error) {
	grpcConn, errGrpc := grpc.NewClient(cfg.DSNgrpc())
	if errGrpc != nil {
		return nil, errGrpc
	}
	grpClient := grpc_gen.NewAuthServiceClient(grpcConn)

	return &GrpcClient{
		Conn:   grpcConn,
		Client: grpClient,
	}, nil
}

// HTTP server to handler REST requests
func RunServer(router *echo.Echo, ctx *context.Context, cfg *config.ServerConfig) error {
	server := echo.StartConfig{
		Address:         cfg.HttpServerPort,
		GracefulTimeout: 5 * time.Second,
	}

	return server.Start(*ctx, router)
}
