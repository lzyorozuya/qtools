package main

import (
	"context"
	"fmt"
	"github.com/lz01wcy/qtools/qcertificate"
	"github.com/lz01wcy/qtools/qgrpc"
	"github.com/lz01wcy/qtools/qgrpc/examples/pb"
)

func main() {
	clientConn, err := qgrpc.NewConn(&qgrpc.ClientConfig{
		RemoteAddr: "127.0.0.1:3456",
		Cert: qcertificate.Certificate{
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
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	client := pb.NewEchoClient(clientConn)
	echoMassage, err := client.Echo(context.Background(), &pb.EchoMassage{Data: "123456789"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(echoMassage.String())
}
