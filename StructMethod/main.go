package main

import "fmt"

type car struct {
	name string
}

func (c car) SetName01(s string) {
	fmt.Printf("1. address: %p\n", &c)
	c.name = s
}

func(c *car) SetName02(s string) {
	fmt.Printf("2. address: %p\n", c)
	c.name = s
}

func main() {
	c := &car{}
	fmt.Printf("c address: %p\n", c)

	c.SetName02("bar")
	fmt.Println(c.name)
	c.SetName01("foo")
	fmt.Println(c.name)
}