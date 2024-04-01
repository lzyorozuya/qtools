package qrandom

import "testing"

func TestHexString(t *testing.T) {
	hexString, err := HexString(16)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(hexString, len(hexString))
}
