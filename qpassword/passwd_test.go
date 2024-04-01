package qpassword

import "testing"

func TestPasswordEncode(t *testing.T) {
	t.Log(Encode("Sxywb_db_2021!"))
	t.Log(Encode(""))
}

func TestPasswordDecode(t *testing.T) {
	t.Log(Decode("aHR2Hv4Dt6xec+7OIR0hVhI1IujvYxVPV0xNIV8X"))
	t.Log(Decode("aHR2Hv4Dt6xec+7OIR0haHR2Hv4Dt6xec+7OIR0h"))
}

func TestDecodeFile(t *testing.T) {
	file, err := DecodeFile("../x509/encode_crt")
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%s\n", file)
}
