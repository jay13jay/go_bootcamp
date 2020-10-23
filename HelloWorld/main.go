package main

import (
	"fmt"
	"time"

	"github.com/jay13jay/go_bootcamp/newlib"
)

func sayHello(name string) {
	fmt.Printf("Hello, %v\n", name)
}

func main() {
	var (
		now time.Time
	)
	now = time.Now()
	sayHello("Josh")
	fmt.Printf("the current time is:\t %v\n", now)

	// Use imported librarys
	newlib.Name()
	fmt.Println("\n System Information")
	fmt.Printf("Go version:\t%v\n", newlib.Goversion())
	fmt.Printf("# of CPU's:\t%d\n", getcpus())
}
