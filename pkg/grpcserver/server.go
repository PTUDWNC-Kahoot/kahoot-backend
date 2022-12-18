package grpcserver

import (
	"net"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	server   *grpc.Server
	listener net.Listener
	notify   chan error
}

func (g *GRPCServer) Start() {
	g.start()
}

func (g *GRPCServer) start() {
	go func() {
		g.notify <- g.server.Serve(g.listener)
		close(g.notify)
	}()
}

func (g *GRPCServer) Notify() <-chan error {
	return g.notify
}

func (g *GRPCServer) Shutdown() {
	g.server.GracefulStop()
}

func New(server *grpc.Server, address string) *GRPCServer {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	return &GRPCServer{
		server:   server,
		listener: lis,
	}
}
