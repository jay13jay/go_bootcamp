package newlib

import (
	"fmt"
	"runtime"
)

func Name() {
	fmt.Println("This function is part of the newlib library")
}

func Goversion() string {
	return runtime.Version()
}
