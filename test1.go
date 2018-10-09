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

	//fmt.Println(num)
	//多变量赋值
	//a, b, c := 1, 2, 3

	/*fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)*/

	var a string
	a = "bear"
	fmt.Println(a)

	var returndata bool
	returndata = false

	fmt.Println(returndata)

	var mycash float32
	//可定义 但不赋值
	mycash = 188.65
	fmt.Println(mycash)

	str1 := "my wife is " //申明并赋值
	str2 := "miss zhang"

	str := str1 + str2 //已经定义了新的变量

	str = str1 + "miss Wang" //拼接字符串

	fmt.Println(str)

	f1 := float32(3.1415926)
	fmt.Println(f1)

	var arr [2]int //申明
	arr[0] = 66    //表达式
	arr[1] = 77
	fmt.Printf("fist num is %d\n", arr[0])
	fmt.Printf("fist num is %d\n", arr[1])

	//var array = [7]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'} //带有=
	//var aSlice, bSlice []byte                              //slice类型
	//aSlice = array[:2]
	//bSlice = array[2:4]

	//astr
	var map1 map[string]int //索引为string 值为int的map类型
	//map1 := make(map[string]int)

	map1["name"] = 1
	map1["age"] = 29

	fmt.Println(map1["age"])

}
