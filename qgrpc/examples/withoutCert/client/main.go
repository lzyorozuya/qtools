package main

import (
	"context"
	"fmt"
	"github.com/lzyorozuya/qtools/qgrpc"
	"github.com/lzyorozuya/qtools/qgrpc/examples/pb"
)

func main() {
	clientConn, err := qgrpc.NewConnOnH2C("127.0.0.1:3456")
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
