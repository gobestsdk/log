package log

import "testing"

//import "github.com/gobestsdk/log"

func TestPrint(t *testing.T) {
	field := Fields{
		"test": "ok",
		"test2": map[string]interface{}{
			"aname": 1,
			"bname": "b",
		},
	}
	Print(Fields{
		"test": "ok",
	})
	SetSerializationtype("table")
	Info(field)
	SetSerializationtype("yaml")
	Info(field)
}

func TestInfoTable(t *testing.T) {
	SetSerializationtype("table")
	Info(Fields{
		"test": "ok",
		"test2": map[string]interface{}{
			"aname": 1,
			"bname": "b",
		},
	})
}
func TestInfoYAML(t *testing.T) {
	SetSerializationtype("yaml")
	Info(Fields{
		"test": "ok",
	})
}
