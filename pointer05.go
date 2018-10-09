package main

import (
	"fmt"
)

func main() {

	var a int

	var b *int

	var c **int

	a = 10

	b = &a

	c = &b

	fmt.Println(a, b, c, *b, *c, **c)

}
