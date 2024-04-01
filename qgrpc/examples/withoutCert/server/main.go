package main

import (
	"context"
	"fmt"
	"github.com/lz01wcy/qtools/qgrpc"
	"github.com/lz01wcy/qtools/qgrpc/examples/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedEchoServer
}

func (s server) Echo(ctx context.Context, massage *pb.EchoMassage) (*pb.EchoMassage, error) {
	return massage, nil
}

func main() {
	svr := server{}
	err := qgrpc.StartServiceH2C(3456, func(s *grpc.Server) {
		pb.RegisterEchoServer(s, svr)
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}
