package log

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
	"runtime"
	"runtime/debug"
	"strconv"
	"time"
)

const (
	PRINT = 0
	INFO  = 1
	WARN  = 2
	ERROR = 3
	FATAL = 4
)

var (
	logpath           string = "log"
	level             int
	serializationtype string
	writelog          bool = false

	//showtracefile 是否显示代码
	showtime      bool = false
	showtracefile bool = false
)

func SetTime(show bool) {
	showtime = show
}
func SetSerializationtype(t string) {
	serializationtype = t
}
func SetCodefile(show bool) {
	showtracefile = show
}

func Setlogfile(f string) {
	writelog = true
	logpath = f
}
func Setlevel(l int) {
	level = l
}

func console_printjson(l int, arg map[string]interface{}) {

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

func print(l int, arg map[string]interface{}) {
	arg["_"] = time.Now().String()[:19]

	var file string
	var line int
	var ok bool
	if l >= FATAL {
		_, file, line, ok = runtime.Caller(3)
		debug.PrintStack()
	} else {
		_, file, line, ok = runtime.Caller(2)
	}

	if ok {
		//f := runtime.FuncForPC(pc)
		arg["_trace"] = file + ":" + strconv.Itoa(line)
	}

	if l >= level && writelog {
		base_print(arg)
	}
	console_printjson(l, arg)
}

var Warnmsgfunc, Errormsgfunc func(arg map[string]interface{})

func Print(arg map[string]interface{}) {
	print(0, arg)
}

func Info(arg map[string]interface{}) {
	print(1, arg)
}

func Warn(arg map[string]interface{}) {
	print(2, arg)
	if Warnmsgfunc != nil {
		Warnmsgfunc(arg)
	}
}

func Error(arg map[string]interface{}) {
	print(3, arg)
	if Errormsgfunc != nil {
		Errormsgfunc(arg)
	}
}

func Fatal(arg map[string]interface{}) {
	print(4, arg)
}
