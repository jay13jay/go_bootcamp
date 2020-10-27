package main

import (
	"fmt"
	"runtime"
)

var numcpu = runtime.NumCPU()

func main() {
	printCPUuse()
	printMemUsage()

}

func printCPUuse() {
	fmt.Printf("\nNumber of CPUs:\t %d\n", numcpu)
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// m := runtime.ReadMemStats(&runtime.MemStats) // This doesn't work, not sure why yet, but I am way ahead of the lectures so will come back later

	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n\n", m.NumGC)
	fmt.Printf("unformated print:\n\t%v", m)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
