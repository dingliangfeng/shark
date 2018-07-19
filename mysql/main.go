package main

import (
	"fmt"
	"github.com/dingliangfeng/mysql/base"
)

type Main struct {
	str string
	base.Base
	Mysql base.Mysql
}

func main() {
	fmt.Println("begin")
	main := Main{}
	db := main.Mysql.Connect()
	fmt.Println(main.str)
	db.Query("INSERT INTO user (name) VALUES ('dingliangfeng')")
	defer db.Close()
}
