package qhash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"github.com/lz01wcy/qtools/qpool"
	"golang.org/x/crypto/sha3"
)

var md5pool = qpool.NewQPool(md5.New)

func PoolMD5(data ...[]byte) []byte {
	h := md5pool.Get()
	h.Reset()
	for _, d := range data {
		h.Write(d)
	}
	result := h.Sum(nil)
	md5pool.Put(h)
	return result
}

func PoolMD5256String(s string) []byte {
	return PoolMD5([]byte(s))
}

func PoolMD5HexString(s string) string {
	return hex.EncodeToString(PoolMD5([]byte(s)))
}

func PoolMD5Base64String(s string) string {
	return base64.StdEncoding.EncodeToString(PoolMD5([]byte(s)))
}

var sha1pool = qpool.NewQPool(sha1.New)

func PoolSHA1(data ...[]byte) []byte {
	h := sha1pool.Get()
	h.Reset()
	for _, d := range data {
		h.Write(d)
	}
	result := h.Sum(nil)
	sha1pool.Put(h)
	return result
}

func PoolSHA1256String(s string) []byte {
	return PoolSHA1([]byte(s))
}

func PoolSHA1HexString(s string) string {
	return hex.EncodeToString(PoolSHA1([]byte(s)))
}

func PoolSHA1Base64String(s string) string {
	return base64.StdEncoding.EncodeToString(PoolSHA1([]byte(s)))
}

var sha3256pool = qpool.NewQPool(sha3.New256)

func PoolSHA3256(data ...[]byte) []byte {
	h := sha3256pool.Get()
	h.Reset()
	for _, d := range data {
		h.Write(d)
	}
	result := h.Sum(nil)
	sha3256pool.Put(h)
	return result
}

func SHA3256String(s string) []byte {
	return PoolSHA3256([]byte(s))
}

func SHA3256HexString(s string) string {
	return hex.EncodeToString(PoolSHA3256([]byte(s)))
}

func SHA3256Base64String(s string) string {
	return base64.StdEncoding.EncodeToString(PoolSHA3256([]byte(s)))
}

var sha256pool = qpool.NewQPool(sha256.New)

func PoolSHA256(data ...[]byte) []byte {
	h := sha256pool.Get()
	h.Reset()
	for _, d := range data {
		h.Write(d)
	}
	result := h.Sum(nil)
	sha256pool.Put(h)
	return result
}

func SHA256256String(s string) []byte {
	return PoolSHA256([]byte(s))
}

func SHA256HexString(s string) string {
	return hex.EncodeToString(PoolSHA256([]byte(s)))
}

func SHA256Base64String(s string) string {
	return base64.StdEncoding.EncodeToString(PoolSHA256([]byte(s)))
}
