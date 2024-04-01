package qrandom

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func Bytes(length int) ([]byte, error) {
	if length <= 0 {
		return nil, fmt.Errorf("不合理的长度")
	}
	buf := make([]byte, length)
	_, err := rand.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func HexString(length int) (string, error) {
	bytes, err := Bytes(length)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), err
}
