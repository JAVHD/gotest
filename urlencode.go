package main

import (
	"fmt"
	"net/url"
)

func main() {
	//param := "{\"password\":\"zhanglanlu\",\"username\":\"wy\"}"
	param := `{"password":"zhanglanlu","username":"wy"}`
	fmt.Println(param)
	//link := "http://link.haiyuyunce.com/lz/?param=" + param
	//link := url.QueryEscape("http://link.haiyuyunce.com/lz/?param=" + param)
	link := "http://link.haiyuyunce.com/lz/login.action?params=" + url.QueryEscape(param)
	fmt.Println(link)

}
