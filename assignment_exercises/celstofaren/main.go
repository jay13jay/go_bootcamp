package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
)

func main() {
	_, progname := path.Split(os.Args[0])
	fmt.Printf("Args provided:\t%v\n", len(os.Args))
	if len(os.Args) < 2 {
		fmt.Printf("Please provide an argument\n\n")
		fmt.Printf("Usage:\n")
		fmt.Printf("%v [Temperature in Celsius]\n\n", progname)
		os.Exit(1)
	}
	temp, _ := strconv.ParseFloat(os.Args[1], 32)
	fmt.Printf("Temperature in C:\t%.2f\n", temp)
	ftemp := temp*9/5 + 32
	fmt.Printf("Temperature in F:\t%.2f\n", ftemp)
}
