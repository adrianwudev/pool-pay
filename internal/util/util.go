package util

import (
	"fmt"
	"runtime"
)

func PrintLineNumber() {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("Line number: %d, File: %s\n", line, file)
}
