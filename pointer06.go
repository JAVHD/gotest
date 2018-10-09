package main

import (
	"fmt"
)

func Foo(a, b int, c *int) {
	*c = a * b
}

func main() {
	a := 10
	b := 5

	c := &a

	fmt.Println(a, b, c)

	Foo(a, b, c)

	fmt.Println(a, b, c)

	//Foo(a, b, *c)

	fmt.Println(a, b, c)

}
