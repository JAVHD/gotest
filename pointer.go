package main

import (
	"fmt"
)

func main() {

	//x := 10

	//var p *int = &x

	//*p += 20

	p := 10

	*p = &p

	fmt.Println(p, &p, *p)

}
