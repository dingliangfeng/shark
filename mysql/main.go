package main

import (
	"github.com/dingliangfeng/mysql/base"
)

type Main struct {
	str string
	base.Base
	Mysql base.Mysql
}

func main() {
	main := Main{}
	main.Mysql.Connect()
	db := main.Mysql.Db()
	db.Query("INSERT INTO user (name) VALUES ('dingliangfeng')")
	defer db.Close()
}
