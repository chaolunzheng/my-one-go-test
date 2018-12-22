package main

import (
	"database/sql"
	"fmt"
)

//连接数据库实例
func main() {
	dirver := "jdbc.mysql.go"
	url := "jdbc:mysql:3306//hive"
	db, e := sql.Open(dirver, url)
	fmt.Println(db, e)
	stmt, _ := db.Prepare("select * from tableNmae where id=10")
	result, i := stmt.Exec()
	for i != nil {
		fmt.Println(result)
	}
}
