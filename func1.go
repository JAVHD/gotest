package test

import ( //import 用于导入包文件
	f "fmt"
)

func add(a *int) int {
	*a = *a + 1
	return *a
}

func init() {
	f.Println("This is init!") //先执行init函数
}

func main() {

	a1 := 1

	a2 := add(&a1)

	f.Println(a2)
	f.Println(a1)

	/*for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}*/

	type person struct { //结构体的申明
		name string
		age  int
	}

	var P person

	P.age = 26

	f.Println("My age is ", P.age)
}
