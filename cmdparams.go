package main

import (
	"flag"
	"fmt"
)

func main() {
	s := flag.String("name", "World", "the name to print")
	flag.Parse()
	fmt.Printf("Hello, %s!", *s)
}
