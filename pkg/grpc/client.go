package grpc

import (
	"github.com/isOdin-l/TinderArt/pkg/configs"
	"github.com/isOdin-l/TinderArt/pkg/grpc/auth"
	grpc "google.golang.org/grpc"
)

// Creation of gRPC client to connect with Auth service
type GrpcClient struct {
	Conn *grpc.ClientConn
	auth.AuthServiceClient
}

func NewGrpcAuthClient(cfg *configs.ConfigGrpcClient) (*GrpcClient, error) {
	grpcConn, errGrpc := grpc.NewClient(cfg.DSN())
	if errGrpc != nil {
		return nil, errGrpc
	}
	grpClient := auth.NewAuthServiceClient(grpcConn)

	return &GrpcClient{
		Conn:              grpcConn,
		AuthServiceClient: grpClient,
	}, nil
}
