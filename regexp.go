package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"regexp"
	//"strings"
	//"reflect"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	//ioutil.ReadAll(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)

	//body = string(body)
	str := string(body) //golang是强类型语言 此处不能强行复制 body []byte

	fmt.Println(str)
	//fmt.Println(reflect.TypeOf(str))

}
