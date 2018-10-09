package test

import (
	//"fmt"
	"io"
	"log"
	"net/http"
)

func HelloAction(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, "Golang Web Server")
	io.WriteString(w, r.Method)
	//io.WriteString(w, "xxx")
}

func main() {
	http.HandleFunc("/HelloWorld", HelloAction) //定义路由
	//监听端口
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("fatal:", err)
	}
}
