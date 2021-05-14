package log

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
)

func console_printmap(l int, arg map[string]interface{}) {

	switch serializationtype {
	case "json":
		fallthrough
	case "yaml":
		console_printjson(l, arg)
	case "table":
		console_table(arg)
	}
}

func console_printstruct(l int, arg Struct) {

	switch serializationtype {
	case "json":
		fallthrough
	case "yaml":
		//console_printjson(l, arg)
	case "table":
		//console_table(arg)
	}
}

func console_printjson(l int, arg Fields) {

	var c Colortext
	var pre bool
	switch l {
	case PRINT:
		c = Black
	case INFO:
		c = Pink
	case WARN:
		c = Yellow
	case ERROR:
		c = LightRed
	case FATAL:
		c = Red
	}
	s := string(c)
	if showtime || l >= WARN {
		s = s + arg["_"].(string)
		pre = true
	}
	if showtracefile || l >= WARN {
		if t, exist := arg["_trace"]; !exist || t == nil {
		} else {
			s = s + " " + arg["_trace"].(string)
			pre = true
		}
	}

	delete(arg, "_")
	delete(arg, "_trace")
	var bs []byte
	var err error
	switch serializationtype {
	case "yaml":
		bs, err = yaml.Marshal(arg)
	default:
		bs, err = json.Marshal(arg)

	}

	if err != nil {
		println(err.Error())
	}
	if pre {
		s = s + "\t" + string(bs) + string(EndColor)
	} else {
		s = s + string(bs) + string(EndColor)
	}

	fmt.Println(s)

}

func console_table(arg map[string]interface{}) {
	bs, err := Fields(arg).TableMarshal()
	if err != nil {
		println(err.Error())
	}
	fmt.Println(string(bs))
}
