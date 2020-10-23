package main

import "fmt"

func sayHello(name string, age int) {
	fmt.Println("Hello,", name)
	fmt.Printf("This programs says you are %v years old\n", age)
}

func main() {
	sayHello("Josh", 30)
}
