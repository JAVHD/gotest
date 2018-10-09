package test

import "fmt"

func main() {

	//定义一个切片值
	aslice := []int{2, 4, 6}
	//sum_num := 0

	//range golang内置函数
	for index, num := range aslice { //index for循环结束后会被销毁
		fmt.Println(index)
		fmt.Println(num)
	}

	fmt.Println(index)

	//fmt.Println(k)
	//fmt.Println(sum_num)

	/*nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)*/
}
