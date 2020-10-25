package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	// fmt.Printf("%#v\n", os.Args)
	// fmt.Printf("%v\n\n", os.Args)

	_, file := path.Split(os.Args[0])

	fmt.Printf("Program\t\t: %v\n", file)
	fmt.Println("Provided Args\t: ", len(os.Args)-1)

	count := 1
	if len(os.Args) >= 1 {
		for i := 0; i < len(os.Args)-1; i++ {
			fmt.Printf("Argument %v\t: %v\n", count, os.Args[count])
			count = count + 1
		}
	}

}
