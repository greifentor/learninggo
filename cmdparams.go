package main

import (
	"flag"
	"fmt"
)

func main() {
	s := flag.String("name", "World", "the name to print")
	flag.Parse()
	fmt.Printf("\nHello, %s!\n", *s)
	fmt.Printf("args: %s\n\n", flag.Args())
	flag.Usage()
	fmt.Println()
	fmt.Println()
}
