package main

import (
	"fmt"
)

type StructType struct {
	name string
	age  int
}

func (s StructType) MyMethod1() {
	fmt.Println("this is MyMethod1")
	//s.name = "bear01"
}

func (s *StructType) MyMethod2() {
	fmt.Println("this is MyMethod2")
	s.name = "bear01"

}

func main() {
	var st = StructType{"bear", 18}

	st.MyMethod1()
	st.MyMethod2()

	fmt.Println(st)
}
