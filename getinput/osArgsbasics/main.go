package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", os.Args)
	fmt.Printf("%v\n\n", os.Args)

	fmt.Printf("Program:\t%v\n", os.Args[0])
	fmt.Println("# of Arguments:\t", len(os.Args))

	count := 1
	if len(os.Args) >= 1 {
		for i := 0; i < len(os.Args)-1; i++ {
			fmt.Printf("Argument %v:\t%v\n", count, os.Args[count])
			count = count + 1
		}
	}

}
