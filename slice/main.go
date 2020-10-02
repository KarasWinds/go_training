package main

import(
	"fmt"
)

func modify(foo []string) {
	foo[1] = "c"
	fmt.Println("modify foo", foo)
}

func addValue(foo []string) []string {
	foo = append(foo, "c")
	fmt.Println("add foo", foo)
	return foo
}

func main() {
	foo := []string{"a", "b"}
	fmt.Println("before modify:", foo)
	modify(foo)
	fmt.Println("after modify:", foo)
	fmt.Println("before add:", foo)
	addValue(foo)
	fmt.Println("after add:", foo)
	foo = addValue(foo)
	fmt.Println("after add:", foo)
	bar := foo[:1]
	fmt.Println("bar:", bar)
	s1 := append(bar, "b")
	fmt.Println("foo:", foo)
	fmt.Println("s1:", s1)
	s2 := append(bar, "c", "d", "f")
	fmt.Println("foo:", foo)
	fmt.Println("s2:", s2)
}