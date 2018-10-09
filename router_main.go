package main

import (
	"fmt"
	"log"
	"net/http"
	"testPkg"
)

func Action(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "another...")

}

func main() {

	//testPkg.AddRouter("/")
	testPkg.AddRouter("/test01")
	testPkg.AddRouter("/test02")
	testPkg.AddRouter("/test03")
	testPkg.AddRouter("/test04")

	//http.HandleFunc("/", Action)

	//监听端口
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		//fmt.Println("ListenAndServe ", err)
		log.Fatal("ListenAndServe: ", err)

	}

	//http.HandleFunc("/", Action)

}
