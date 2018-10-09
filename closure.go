package main

import (
	"fmt"
)

func ClosureTest(x int) func(int) int {

	return func(y int) int {
		return x + y
	}

}

func main() {

	a := ClosureTest(10)

	fmt.Println(a(100))
}
