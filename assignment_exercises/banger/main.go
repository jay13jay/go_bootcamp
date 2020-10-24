package main

import (
	"fmt"
	"os"
	"strings"
)

var word string

func printbang(word string) {
	banger := strings.Repeat("!", len(word))
	word = strings.ToUpper(word)
	// fmt.Printf("%s%s\n", word, banger)
	fmt.Printf("%s%s%s\n", banger, word, banger)

}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("no word provided!")
		os.Exit(1)
	}
	word = os.Args[1]

	printbang(word)

}
