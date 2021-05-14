package log

import (
	"encoding/json"
	"github.com/gobestsdk/types"
	"runtime"
	"runtime/debug"
	"strconv"
	"time"
)

func printfixedmap(l int, arg map[string]interface{}) {
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
		a, _ := json.Marshal(arg)
		base_print(string(a))
	}
	console_printmap(l, arg)
}

func printfixedstruct(l int, arg Struct) {
	arg = append(arg, KV{"_", time.Now().Format(types.CommonDatetime)})
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
		arg = append(arg, KV{"_trace", file + ":" + strconv.Itoa(line)})
	}

	if l >= level && writelog {
		a, _ := json.Marshal(arg)
		base_print(string(a))
	}
	console_printstruct(l, arg)
}
