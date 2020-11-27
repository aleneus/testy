package testy

import (
	"fmt"
	"runtime"
)

func info() (s string) {
	var lineSearched int
	var fileSearched string

	for i := 0; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		fun := runtime.FuncForPC(pc)
		if fun == nil {
			break
		}

		if fun.Name() == "testing.tRunner" {
			break
		}

		lineSearched = line
		fileSearched = file
	}

	return fmt.Sprintf("%s:%d:", fileSearched, lineSearched)
}
