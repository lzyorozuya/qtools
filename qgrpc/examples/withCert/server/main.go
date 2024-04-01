package main

import (
	"context"
	"fmt"
	"github.com/lzyorozuya/qtools/qcertificate"
	"github.com/lzyorozuya/qtools/qgrpc"
	"github.com/lzyorozuya/qtools/qgrpc/examples/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedEchoServer
}

func (s server) Echo(ctx context.Context, massage *pb.EchoMassage) (*pb.EchoMassage, error) {
	fmt.Println("收到消息: ", massage.String())
	return massage, nil
}

func main() {
	config := qgrpc.ServerConfig{
		Port: 3456,
		Certs: qcertificate.Certs{
			{
				Name: "JiangNan7Guai",
				Crt: qcertificate.EncodedFile{
					Path:      "qgrpc/examples/withCert/x509/encode_crt",
					IsEncoded: true,
				},
				Key: qcertificate.EncodedFile{
					Path:      "qgrpc/examples/withCert/x509/encode_key",
					IsEncoded: true,
				},
			},
		},
	}
	svr := server{}
	err := qgrpc.StartService(config, func(s *grpc.Server) {
		pb.RegisterEchoServer(s, svr)
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}
