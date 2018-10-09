/*
* @Author: bear
* @Date:   2018-09-03 11:42:59
* @Last Modified by:   bear
* @Last Modified time: 2018-09-03 21:32:48
 */

package main

import (
	"fmt"
)

func test(x int) (func() int, func()) {
	return func() int {
			x += 100
			return x + 100
		}, func() {
			fmt.Println(x)
		}
}

func main() {
	a, b := test(10)

	a()
	b()

}
