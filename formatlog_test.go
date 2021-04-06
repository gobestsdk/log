package log

import "testing"

func TestPrint(t *testing.T) {
	Print(Fields{
		"test": "ok",
	})
}
