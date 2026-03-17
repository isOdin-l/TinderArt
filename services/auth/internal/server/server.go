package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/isOdin-l/TinderArt/pkg/configs"
	grpc_auth "github.com/isOdin-l/TinderArt/pkg/grpc/auth"
	config "github.com/isOdin-l/TinderArt/services/auth/configs"
	"github.com/labstack/echo/v5"
	"google.golang.org/grpc"
)

type IHandlerGrpc interface {
	grpc_auth.AuthServiceServer
}

type Server struct {
	httpServerCfg *echo.StartConfig
	httpRouter    *echo.Echo

	grpcServer *grpc.Server
	listener   net.Listener
}

func NewServer(cfgHTTP *config.ServerConfig, cfgGRPC *configs.ConfigGrpcServer, routerHTTP *echo.Echo, routerGRPC grpc_auth.AuthServiceServer) (*Server, error) {
	var server Server

	// HTTP
	server.httpServerCfg = &echo.StartConfig{
		Address:         cfgHTTP.HttpServerPort,
		GracefulTimeout: 5 * time.Second,
	}

	// GRPC
	listen, errListn := net.Listen("tcp", cfgGRPC.DSN())
	if errListn != nil {
		return nil, errListn
	}

	server.listener = listen
	server.grpcServer = grpc.NewServer()
	grpc_auth.RegisterAuthServiceServer(server.grpcServer, routerGRPC)

	return &server, nil
}

func (server *Server) RunServer() {
	// Definition context for server's graceful shutdown
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	ch := make(chan error)

	go server.runGRPC(ctx, ch)
	go server.runHTTP(ctx, ch)

	// если ошибка, то читаем из канала и закрываем его, чтобы освободить grpc http горутины
	// Если выход по контексту, то горутины сами освободятся, потому что мы им прокинули контекст
	var errSrv error
	select {
	case <-ctx.Done():
	case errSrv = <-ch:
		close(ch)
		server.httpRouter.Logger.Error(fmt.Sprintf("Server error: %s", errSrv.Error()))
	}

}

// hadnler GRPC server
func (server *Server) runGRPC(ctx context.Context, ch chan error) {
	select {
	case ch <- server.grpcServer.Serve(server.listener):
	case <-ctx.Done():
	}

	server.grpcServer.GracefulStop()
}

// HTTP server to handler REST requests
func (server *Server) runHTTP(ctx context.Context, ch chan error) {
	select {
	case ch <- server.httpServerCfg.Start(ctx, server.httpRouter):
	case <-ctx.Done():
	}
}
