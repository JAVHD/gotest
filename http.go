package main

import (
	"fmt"
	"io"
	"net/http"
	//"reflect"
)

func MyServer(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, "this is a web server!")
	//io.WriteString(w, r.Method)
	//io.WriteString(w, r.Method)
	//io.WriteString(w, "xxx")
	r.ParseForm()
	//io.WriteString(w, reflect.TypeOf(r.Form))

	apiData := make(map[string]interface{})
	//apiData["password"] = r.Form("password")
	//apiData["password"] = r.PostFormValue("password")
	//apiData["password"] = 123456

	fmt.Fprint(w, r.Form)
	fmt.Fprint(w, apiData)
	//fmt.Fprint(w, r.Form["apiData"])

	//io.WriteString(w, r.Form.Get("data"))

}

func GetMethod(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.Method)
}

func main() {
	http.HandleFunc("/index", MyServer)
	http.HandleFunc("/getMethod", GetMethod)

	//启动服务 绑定端口
	http.ListenAndServe(":8081", nil)

}
