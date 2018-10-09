package main

import (
	"fmt"
)

// a := (func() func() {

// 	return func() {
// 		fmt.Println(123)
// 	}

// })()

var j int = 5

b := func() func(){
       //var i int = 10
		i := 10

       return func(){
       		fmt.Printf("i,j:%d,%d\n",i,j)
       }
     }()

func main() {

	b()

}