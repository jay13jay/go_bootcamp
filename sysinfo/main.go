package main

import (
	"fmt"
	"runtime"
)

func main() {
	var num_cpu = runtime.NumCPU()
	fmt.Printf("Number of CPUs:\t %d\n", num_cpu)
}
