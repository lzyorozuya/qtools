package qcrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/lzyorozuya/qtools/qhash"
)

func AESGCMEncrypt(key []byte, data []byte) ([]byte, error) {
	var err error
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("很严重的错误呢,差点duang了%s", e)
		}
	}()
	hashedKey := qhash.PoolSHA256(key)
	block, err := aes.NewCipher(hashedKey)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return gcm.Seal(nil, hashedKey[:gcm.NonceSize()], data, nil), nil
}

func AESGCMDecrypt(key []byte, date []byte) (_ []byte, err error) {
	hashedKey := qhash.PoolSHA256(key)
	block, err := aes.NewCipher(hashedKey)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return gcm.Open(nil, hashedKey[:gcm.NonceSize()], date, nil)
}

type CodeHandler func([]byte, []byte) ([]byte, error)

func EncodeBase64(key, data []byte, handler CodeHandler) (string, error) {
	bytes, err := handler(key, data)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), err
}

func DecodeBase64(key, base64Str string, handler CodeHandler) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil, err
	}
	bytes, err := handler([]byte(key), data)
	if err != nil {
		return nil, err
	}
	return bytes, err
}

func EncodeHex(key, data []byte, handler CodeHandler) (string, error) {
	bytes, err := handler(key, data)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), err
}

func DecodeHex(key, base64Str string, handler CodeHandler) ([]byte, error) {
	data, err := hex.DecodeString(base64Str)
	if err != nil {
		return nil, err
	}
	bytes, err := handler([]byte(key), data)
	if err != nil {
		return nil, err
	}
	return bytes, err
}
