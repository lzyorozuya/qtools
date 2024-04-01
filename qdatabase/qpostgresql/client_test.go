package qpostgresql

import "testing"

func TestNewClient(t *testing.T) {
	info := &Info{
		Host:            "localhost",
		Port:            "5432",
		Account:         "pg",
		PasswordDecoded: "",
		Password:        "pass",
		DatabaseName:    "crud",
	}
	_, err := NewClient(info, []any{})
	if err != nil {
		t.Log(err)
		return
	}

}
