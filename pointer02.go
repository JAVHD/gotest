package main

import (
	"fmt"
)

func main() {

	a := 10

	b := &a

	*b++

	fmt.Println(a, b)

}
