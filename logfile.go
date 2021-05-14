package log

import (
	"os"
)

func WriteLog(w bool) {
	writelog = w
}
func makefile() {
	_, err := os.Stat(logpath)
	if err != nil {
		console_printmap(INFO, Fields{"log": "create log file", "filename": logpath})
		fs, err := os.OpenFile(logpath, os.O_CREATE|os.O_RDWR, os.ModePerm)
		if err != nil {
			console_printmap(FATAL, Fields{"err": err})
			return
		}
		fs.Close()
	}
}
