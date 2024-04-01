package qcertificate

import (
	"crypto/tls"
	"github.com/lzyorozuya/qtools/qpassword"
	"os"
)

type EncodedFile struct {
	Path      string
	IsEncoded bool
}

func (f EncodedFile) DecodeBytes() ([]byte, error) {
	file, err := os.ReadFile(f.Path)
	if err != nil {
		return nil, err
	}
	if !f.IsEncoded {
		return file, nil
	}
	return qpassword.DecodeBytes(string(file))
}

type Certificate struct {
	Name string      //签名的名称
	Crt  EncodedFile //crt文件
	Key  EncodedFile //key文件
}

func (p Certificate) Cert() (tls.Certificate, error) {
	crt, err := p.Crt.DecodeBytes()
	if err != nil {
		return tls.Certificate{}, err
	}
	key, err := p.Key.DecodeBytes()
	if err != nil {
		return tls.Certificate{}, err
	}
	return tls.X509KeyPair(crt, key)
}

type Certs []Certificate

func (c Certs) Certs() ([]tls.Certificate, error) {
	var res []tls.Certificate
	for _, pair := range c {
		cert, err := pair.Cert()
		if err != nil {
			return nil, err
		}
		res = append(res, cert)
	}
	return res, nil
}
