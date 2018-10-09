/*
* @Author: bear
* @Date:   2018-09-03 22:24:39
* @Last Modified by:   bear
* @Last Modified time: 2018-09-03 23:09:13
 */

package main

import "fmt"

func main() {
	resule := getNum()
	fmt.Println(resule())
}
func getNum() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// func main() {
// 	var f = Adder()
// 	fmt.Println(f(1), "-")
// 	fmt.Println(f(20), "-")
// 	fmt.Println(f(300), "-")

// }
// func Adder() func(int) int {
// 	var x int
// 	return func(delta int) int {

// 		fmt.Println(x)

// 		x += delta
// 		return x
// 	}
// }
