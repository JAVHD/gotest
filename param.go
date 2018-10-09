package main

import (
	"fmt"
)

//变参数 只能放在尾部
func MyFunc(a string, b ...int) {
	fmt.Printf("%T %v\n", b, b)
}

func MyFuncString(a ...string) {
	fmt.Printf("%T %v\n", a, a)
}

func MyFuncInterface(a ...interface{}) {
	fmt.Printf("%T %v\n", a, a)
}

func main() {
	MyFunc("bear", 10)
	MyFuncString("bear", "yong", "zhao")
	MyFuncInterface(18, "bear", 58.25, []int{1, 2, 3})
}
