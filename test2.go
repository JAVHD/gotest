package test

import (
	"fmt"
)

func main() {
	//fmt.Println("My first Go code!");
	//var num type num int8

	num := 1

	for i := 0; i < 10; i++ {
		num = num + 1
	}

	fmt.Println(num)

	fmt.Println("hello world")
}
