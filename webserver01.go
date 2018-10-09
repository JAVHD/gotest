package test

import (
	"fmt"
	"log"
	"net/http"
)

func webAction(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.Scheme)

	//fmt.Fprintf(w, "hello world")

	fmt.Fprintf(w, "Hello Wrold!")

}

func main() {

	//设置路由
	http.HandleFunc("/hello", webAction)

	//绑定端口
	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}

}
