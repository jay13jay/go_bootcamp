package main

import (
	"runtime"
)

func getcpus() int {
	return runtime.NumCPU()
}
