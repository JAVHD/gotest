package main //当前程序的包名称

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
)

func main() {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte("123456"))

	pwd := md5Ctx.Sum(nil)
	//pwdStr := string(pwd)
	//fmt.Println(pwdStr)
	fmt.Println(reflect.TypeOf(pwd))
	fmt.Println(hex.EncodeToString(pwd))
}
