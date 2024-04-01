package qgrpc

import (
	"crypto/tls"
	"fmt"
	"github.com/lzyorozuya/qtools/qcertificate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"net"
)

type ServerConfig struct {
	Port  int
	Certs qcertificate.Certs
}

func (g *ServerConfig) Check() error {
	if len(g.Certs) == 0 {
		return fmt.Errorf("没有有效的x509证书")
	}
	return nil
}

func StartService(conf ServerConfig, registerOpt func(s *grpc.Server)) error {
	certs, err := conf.Certs.Certs()
	if err != nil {
		return fmt.Errorf("x509证书解码失败: %s", err)
	}
	if len(certs) == 0 {
		return fmt.Errorf("x509证书解码失败: 没有有效的证书")
	}
	s := grpc.NewServer(
		grpc.Creds(
			credentials.NewTLS(
				&tls.Config{
					Certificates: certs,
					MinVersion:   tls.VersionTLS13,
					CipherSuites: []uint16{
						tls.TLS_AES_256_GCM_SHA384,
						tls.TLS_CHACHA20_POLY1305_SHA256,
					},
				},
			),
		),
	)

	registerOpt(s)

	listener, err := net.Listen("tcp4", fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		return err
	}
	if err = s.Serve(listener); err != nil {
		return fmt.Errorf("http监听失败:%s", err.Error())
	}
	return nil
}

func StartServiceH2C(port int, registerOpt func(s *grpc.Server)) error {
	s := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
	)

	registerOpt(s)

	listener, err := net.Listen("tcp4", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	if err = s.Serve(listener); err != nil {
		return fmt.Errorf("http监听失败:%s", err.Error())
	}
	return nil
}
