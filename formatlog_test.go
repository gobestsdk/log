package log

import "testing"

func TestPrint(t *testing.T) {
	Print(Fields{
		"test": "ok",
	})
}

func TestInfo(t *testing.T) {
	SetSerializationtype("table")
	Info(Fields{
		"test": "ok",
	})
}
