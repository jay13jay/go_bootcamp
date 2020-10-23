package main

import (
	"fmt"

	"github.com/jay13jay/go_bootcamp/testlibrary"
)

func sayHello(name string, age int) {
	fmt.Printf("Hello, %v", name)
	fmt.Printf("\nThis programs says you are %v years old\n", age)
}

func main() {
	sayHello("Josh", 30)

	// Use imported library
	testlibrary.Name()
}
