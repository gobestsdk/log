package log

import (
	"github.com/gobestsdk/trace"

	"os"
	"sync"
	"sync/atomic"
	"time"
)

type buffer struct {
	Data      [1000]string
	idxpr     *int32
	Pre, Next *buffer
}

var (
	pool = sync.Pool{
		New: func() interface{} {
			var i int32 = 0
			return &buffer{
				idxpr: &i,
			}
		},
	}
	buffersw, buffersr *buffer
	buffer_lock        sync.Mutex
)

func base_print(arg string) {
	if buffersw == nil {
		buffersw = pool.Get().(*buffer)
	}
	if atomic.LoadInt32(buffersw.idxpr) >= 1000 {
		nb := pool.Get().(*buffer)
		nb.Pre = buffersw
		if buffersr == nil {
			buffersr = buffersw
			buffersr.Next = nb
		}
		buffersw = nb
	}

	buffersw.Data[atomic.LoadInt32(buffersw.idxpr)] = arg
	atomic.AddInt32(buffersw.idxpr, 1)

}

var Loggoid uint64

func Syncrite() {
re:
	Loggoid = trace.GetGoroutineID()
	makefile()
	afs, err := os.OpenFile(logpath, os.O_APPEND|os.O_RDWR, os.ModeAppend)
	if err != nil {
		console_printmap(INFO, Fields{"log": "open log file", "err": err})
		return
	}
	Info(Fields{
		"logstate": "restartlog",
	})

	defer afs.Close()
	for Loggoid == trace.GetGoroutineID() {
		if buffersr != nil && buffersr != buffersw {
			buffersr = buffersr.Next
		} else {
			time.Sleep(time.Second)
			buffersr = buffersw
			buffersw = pool.Get().(*buffer)
			buffersw.Pre = buffersw
			buffersw.Next = nil
		}

		for i := int32(0); buffersr != nil && i < *buffersr.idxpr; i++ {
			arg := buffersr.Data[i]
			bs := arg
			afs.WriteString(bs + "\n")
		}
	}
	goto re
}
