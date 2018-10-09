package main

import (
	"fmt"
)

func main() {

	a := [4]int{1, 2, 3, 4}

	//b := 10

	fmt.Println(a)

	var p *[4]int = &a

	//var m = &a

	//var n = &b

	fmt.Println(p, *p)

}
