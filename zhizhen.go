package test

import "fmt"

func main() {
	var i int

	i = 5

	var p *int
	p = &i

	fmt.Println(&i) // 返回变量i的地址
	fmt.Println(p)  // 返回变量i的地址

}
