package main

import (
	"fmt"
)

func main() {

	//申明p为一个指针变量后 就会分配一个内存地址
	var p *string

	fmt.Println(&p, p)

	//p的值为nil 空指针不能进行*p操作 解决方法 给p分配一个指向
	var m string

	p = &m
	*p = "this is ..."

	fmt.Println(*p)
	fmt.Println(m)

}
