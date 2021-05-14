package log

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
	serializationtype string = "json"
	writelog          bool   = false

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

func Print(arg map[string]interface{}) {
	printfixedmap(0, arg)
}

func PrintStruct(arg Struct) {
	printfixedstruct(0, arg)
}

func Info(arg map[string]interface{}) {
	printfixedmap(1, arg)
}

var Warnmsgfunc, Errormsgfunc func(arg map[string]interface{})

func Warn(arg map[string]interface{}) {
	printfixedmap(2, arg)
	if Warnmsgfunc != nil {
		Warnmsgfunc(arg)
	}
}

func Error(arg map[string]interface{}) {
	printfixedmap(3, arg)
	if Errormsgfunc != nil {
		Errormsgfunc(arg)
	}
}

func Fatal(arg map[string]interface{}) {
	printfixedmap(4, arg)
}
