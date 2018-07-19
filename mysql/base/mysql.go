package base

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	user string
	pass string
	db   *sql.DB
}

func (mysql *Mysql) Connect() *sql.DB {
	mysql.user = "shark"
	mysql.pass = "lfding"
	dsn := mysql.user + ":" + mysql.pass + "@/go"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("connect mysql error")
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("connect ok")
	mysql.db = db
	return db
}
