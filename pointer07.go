package main

import (
	"fmt"
)

func Foo(v *int) {

	//x := 456
	//v = &x
	//*v = x
	//v = x
	fmt.Printf("%T", *v)
}

func main() {

	x := 1000

	fmt.Println(x)

	Foo(&x)

	fmt.Println(x)

	var s string
	s = "bear"

	var p *string

	p = &s

	fmt.Println(p, *p)

	fmt.Printf("%T %T", p, *p)

}
