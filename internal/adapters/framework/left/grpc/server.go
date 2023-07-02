package rpc

import (
	"log"
	"net"

	"github.com/ishanshre/GoHex/internal/adapters/framework/left/grpc/pb"
	"github.com/ishanshre/GoHex/internal/ports"
	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.ApiPort
}

func NewAdapater(api ports.ApiPort) *Adapter {
	return &Adapter{
		api: api,
	}
}

func (grpca Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen in port 9000: %v", err)
	}
	arithmeticServceServer := grpca
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServceServer)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to solved gRPC server over port 9000: %v", err)
	}
}
