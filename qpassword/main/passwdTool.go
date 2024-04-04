package main

import (
	"flag"
	"github.com/lzyorozuya/qtools/qpassword"
	"os"
	"path/filepath"
)

func main() {
	var toEnc, toDec, encPath, decPath string
	flag.StringVar(&toEnc, "enc", "", "需要编码的字符串")
	flag.StringVar(&toDec, "dec", "", "需要解码的字符串")
	flag.StringVar(&encPath, "encFile", "", "需要编码的文件路径")
	flag.StringVar(&decPath, "decFile", "", "需要解码的文件路径")
	flag.Parse()
	if toEnc != "" {
		rs, err := qpassword.Encode(toEnc)
		if err != nil {
			println(err.Error())
		} else {
			println(rs)
		}
		return
	}
	if toDec != "" {
		rs, err := qpassword.Decode(toDec)
		if err != nil {
			println(err.Error())
		} else {
			println(rs)
		}
		return
	}
	if encPath != "" {
		encodeFile(encPath)
		return
	}
	if decPath != "" {
		decodeFile(decPath)
		return
	}
}

func encodeFile(encPath string) {
	encPath, err := filepath.Abs(encPath)
	if err != nil {
		println(err.Error())
		return
	}
	dir, f := filepath.Split(encPath)
	data, err := os.ReadFile(encPath)
	if err != nil {
		println(err.Error())
		return
	}
	encData, err := qpassword.EncodeBytes(data)
	if err != nil {
		println(err.Error())
		return
	}
	newPath := filepath.Join(dir, "encode_"+f)
	encFile, err := os.Create(newPath)
	if err != nil {
		println(err.Error())
		return
	}
	_, err = encFile.WriteString(encData)
	if err != nil {
		println(err.Error())
		return
	}
	println("加密后的文件已写入", newPath)
	return
}

func decodeFile(decPath string) {
	decPath, err := filepath.Abs(decPath)
	if err != nil {
		println(err.Error())
		return
	}
	dir, f := filepath.Split(decPath)
	data, err := os.ReadFile(decPath)
	if err != nil {
		println(err.Error())
		return
	}
	decData, err := qpassword.DecodeBytes(string(data))
	if err != nil {
		println(err.Error())
		return
	}
	newPath := filepath.Join(dir, "decode_"+f)
	decFile, err := os.Create(newPath)
	if err != nil {
		println(err.Error())
		return
	}
	_, err = decFile.Write(decData)
	if err != nil {
		println(err.Error())
		return
	}
	println("解密后的文件已写入", newPath)
	return
}
