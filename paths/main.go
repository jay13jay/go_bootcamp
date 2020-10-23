package main

import (
	"fmt"
	"path"
)

var mypath string = "/Users/jhax/go/src/github.com/jay13jay/go_bootcamp/paths/main.go"

func main() {
	var dir, file string
	dir, file = path.Split(mypath)
	fmt.Printf("Directory:\t%v\n", dir)
	fmt.Printf("Filename:\t%v\n", file)

	fmt.Println(dir[1])
}
