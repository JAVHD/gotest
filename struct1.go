package test

import "fmt"

type School struct {
	name string
	age  int
}

type Class struct {
	School
	name  string
	sex   int
	score int
}

type Rectangle struct { //长方形
	width  float64 //float必须指定32/64
	height float64
}

type Circle struct { //圆
	r float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	const Pi = 3.1415926

	return Pi * c.r * c.r
}

func main() {

	c1 := Circle{3}

	//调用结构体
	class1 := Class{School{"bear001", 28}, "bear002", 1, 98} //字段名的重载

	fmt.Println(class1.name)

	fmt.Println(c1.area())

	fmt.P

}
