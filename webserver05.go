/*
* @Author: bear
* @Date:   2018-09-17 09:59:27
* @Last Modified by:   bear
* @Last Modified time: 2018-09-17 10:27:39
 */

package main

import (
	"fmt"
	"log"
	"net/http"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "hello bear\n")

}

func AddRouter(pattern string) {

	http.HandleFunc(pattern, MyHandler)
}

func main() {

	AddRouter("/bear/getList")
	AddRouter("/bear/getUser")
	AddRouter("/bear/addUser")

	//http.ListenAndServe(addr, handler)

	err := http.ListenAndServe(":8083", nil)

	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
