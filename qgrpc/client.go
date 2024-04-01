package qgrpc

import (
	"crypto/x509"
	"fmt"
	"github.com/lzyorozuya/qtools/qcertificate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientConfig struct {
	RemoteAddr string
	Cert       qcertificate.Certificate
}

func NewConn(conf *ClientConfig) (*grpc.ClientConn, error) {
	cert, err := conf.Cert.Cert()
	if err != nil {
		return nil, err
	}
	if len(cert.Certificate) <= 0 {
		return nil, fmt.Errorf("不能转换为x509证书")
	}
	x509Cert, err := x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		return nil, err
	}
	cp := x509.NewCertPool()
	cp.AddCert(x509Cert)
	crt := credentials.NewClientTLSFromCert(cp, conf.Cert.Name)
	clientConn, err := grpc.Dial(conf.RemoteAddr, grpc.WithTransportCredentials(crt))
	if err != nil {
		return nil, err
	}
	return clientConn, nil
}

func NewConnOnH2C(remoteAddr string) (*grpc.ClientConn, error) {
	clientConn, err := grpc.Dial(remoteAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return clientConn, nil
}
