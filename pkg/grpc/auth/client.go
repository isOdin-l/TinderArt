package __

import (
	"github.com/isOdin-l/TinderArt/pkg/configs"
	grpc "google.golang.org/grpc"
)

// Creation of gRPC client to connect with Auth service
type GrpcClient struct {
	Conn   *grpc.ClientConn
	Client AuthServiceClient
}

func NewGrpcAClient(cfg *configs.ConfigGrpcClient) (*GrpcClient, error) {
	grpcConn, errGrpc := grpc.NewClient(cfg.DSN())
	if errGrpc != nil {
		return nil, errGrpc
	}
	grpClient := NewAuthServiceClient(grpcConn)

	return &GrpcClient{
		Conn:   grpcConn,
		Client: grpClient,
	}, nil
}
