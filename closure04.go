/*
* @Author: bear
* @Date:   2018-09-03 21:15:43
* @Last Modified by:   bear
* @Last Modified time: 2018-09-03 21:53:39
 */

package main

import (
	"fmt"
)

func main() {
	func(x int) int {

		x += 10
		fmt.Println(x)
		return x
	}(100)
}
