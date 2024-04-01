package qlog

import "testing"

func TestError(t *testing.T) {
	Debug("Debug")
	Info("Info")
	Warn("Warn")
	Error("Error")
}
