package main

import (
	"fmt"

	_ "./bar"
	_ "./foo"
)

var global = convert()

func convert() int {
	return 100
}

func init() {
	global = 0
}

func main() {
	fmt.Println("global:", global)
}
