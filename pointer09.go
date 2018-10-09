package main

import (
	"fmt"
)

func MyFunc(x *int) {

	fmt.Println(&x, x)

}

func main() {

	a := 10

	MyFunc(&a)
}
