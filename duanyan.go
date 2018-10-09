package main

import "fmt"

func main() {
	//var m = make(list, 3)

	//type Element interface{} //定义一个接口

	type List interface{} // 定义变量List 为接口类型 interface可以被任何类型所实现

	list := make([]List, 3)
	list[0] = 1
	list[1] = "b"
	list[2] = "c"

	//aa := 'a'
	aa := "abc"
	bb := 'A'

	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Printf("%T", aa)
	//return

	//for index, xxx := range list { //xxx为interface类型
	if value, ok := xxx.(string); ok {
		fmt.Println(index, value)
	}

	//fmt.Println(xxx.(type))

	//}

}
