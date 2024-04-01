package qcrypto

import (
	"encoding/hex"
	"testing"
)

func TestAESGCMEncrypt(t *testing.T) {
	encrypt, err := AESGCMEncrypt([]byte("安忍不动如大地"), []byte("静虑深密如秘藏"))
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%x", encrypt)
}

func TestAESGCMDecrypt(t *testing.T) {
	bytes, err := hex.DecodeString("d291968105cd3579b0a4717af59acb3c2730281c2a84bdff837684934fba74307b16ec4e8f")
	if err != nil {
		return
	}
	decrypt, err := AESGCMDecrypt([]byte("安忍不动如大地"), bytes)
	if err != nil {
		return
	}
	t.Logf("%s", decrypt)
}
