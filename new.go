/*
* @Author: bear
* @Date:   2018-10-08 08:47:45
* @Last Modified by:   bear
* @Last Modified time: 2018-10-08 08:50:12
 */

package main

import (
	"fmt"
)

func main() {

	var i *int

	i = new(int)

	*i = 10

	fmt.Println(i)

}
