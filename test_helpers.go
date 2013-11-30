package greenspun

import (
	"fmt"
	"path"
	"runtime"
	"testing"
)

func ConfirmPanic(t *testing.T, message string, params ...interface{}) func() {
	return func() {
		if recover() == nil {
			t.Fatalf(message, params...)
		}
	}
}

func CallSite() (r string) {
	var stack_trace []uintptr
	runtime.Callers(1, stack_trace)
	my_path := "(unknown)"
	for i, u := range stack_trace {
		if f := runtime.FuncForPC(u); f != nil {
			file, line := f.FileLine(u)
			filepath, filename := path.Split(file)
			if i == 0 {
				my_path = filepath
			} else if my_path != filepath {
				r = fmt.Sprintf("%v:%v", filename, line)
				break
			}
		} else {
			r = "(unknown)"
			continue
		}
	}
	return
}
