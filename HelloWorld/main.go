package main

import (
	"fmt"

	"./helloworld"
	test "./helloworld"
	_ "./helloworld"
)

var (
	foo string = "Hello" //Global
)

const (
	Monday = iota + 1
	Tuesday
)

func main() {
	bar := 123 //local
	fmt.Println(foo)
	fmt.Println(bar)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(helloworld.HelloWorld())
	fmt.Println(test.HelloWorld())
}
