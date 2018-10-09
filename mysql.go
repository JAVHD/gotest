package test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
)

func main() {
	//db, err := sql.Open("mysql", "user:password@tcp(localhost:5555)/dbname?charset=utf8")
	//db 是一个*sql.DB类型的指针，在后面的操作中，都要用到db，err是错误信息，如果为空(nil)说明连接成功，charset设定字符集。

	db, _ := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_test?charset=utf8")

	//checkErr(err)
	fmt.Println(reflect.TypeOf(db))
	return

	stmt, _ := db.Prepare("insert into go_user (username, age) values (?, ?)")
	//checkErr(_)

	res, _ := stmt.Exec("bear", 26)
	//checkErr(_)

	id, _ := res.LastInsertId()
	//checkErr(err)
	//checkTest(err)

	fmt.Println(id)

}
