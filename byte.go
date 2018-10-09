package main

import (
	"fmt"
	//"reflect"
)

func main() {
	//字节转字符串
	fmt.Println(string([]byte{97, 98, 100}))
	//cannot use "z" (type string) as type byte in array or slice literal
	//fmt.Println(string([]byte{"z", 98, 100}))

	fmt.Println(string([]byte{65, 98, 100}))

	//字符串转字节
	fmt.Println([]byte("hello world"))
	fmt.Println([]byte("你好世界"))

	//type error is not an expression
	//fmt.Println(reflect.TypeOf(error))

	fmt.Pr

}
