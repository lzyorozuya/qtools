package qpassword

import (
	"github.com/lz01wcy/qtools/qcrypto"
	"os"
)

var defaultKey = "安忍不动如大地"

func Encode(data string) (string, error) {
	return qcrypto.EncodeBase64(
		[]byte(defaultKey),
		[]byte(data),
		qcrypto.AESGCMEncrypt,
	)
}

func EncodeBytes(data []byte) (string, error) {
	return qcrypto.EncodeBase64([]byte(defaultKey), data, qcrypto.AESGCMEncrypt)
}

func Decode(encoded string) (string, error) {
	data, err := qcrypto.DecodeBase64(defaultKey, encoded, qcrypto.AESGCMDecrypt)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func DecodeBytes(encoded string) ([]byte, error) {
	return qcrypto.DecodeBase64(defaultKey, encoded, qcrypto.AESGCMDecrypt)
}

func DecodeFile(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return qcrypto.DecodeBase64(defaultKey, string(file), qcrypto.AESGCMDecrypt)
}

func EncodeFile(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	base64, err := qcrypto.EncodeBase64([]byte(defaultKey), file, qcrypto.AESGCMDecrypt)
	if err != nil {
		return nil, err
	}
	return []byte(base64), nil
}
