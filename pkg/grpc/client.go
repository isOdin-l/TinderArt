package grpc

import (
	"log"
	"time"

	"github.com/isOdin-l/TinderArt/pkg/configs"
	"github.com/isOdin-l/TinderArt/pkg/grpc/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

// Creation of gRPC client to connect with Auth service
type GrpcClient struct {
	Conn *grpc.ClientConn
	auth.AuthServiceClient
}

const MAX_TRIES = 20

func NewGrpcAuthClient(cfg *configs.ConfigGrpcClient) (*GrpcClient, error) {
	var errGrpc error
	var grpcConn *grpc.ClientConn

	retryableStatusCodes := map[codes.Code]bool{
		codes.Unavailable: true, // etc
	}

	for i := range MAX_TRIES {
		grpcConn, errGrpc = grpc.NewClient(cfg.DSN(), grpc.WithTransportCredentials(insecure.NewCredentials())) // Да, тут надо бы TLS, но я пока хз

		if !retryableStatusCodes[status.Code(errGrpc)] {
			break
		}

		backoff := time.Duration(i+1) * time.Second
		log.Printf("Error calling MyRPC: %v; retrying in %v", errGrpc, backoff)
		time.Sleep(backoff)
	}

	if errGrpc != nil {
		log.Printf("Error calling MyRPC: %v", errGrpc)
		return nil, errGrpc
	}

	return &GrpcClient{
		Conn:              grpcConn,
		AuthServiceClient: auth.NewAuthServiceClient(grpcConn),
	}, nil
}
