package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/open-falcon/hbs/g"
	"log"
)

var DB *sql.DB

func Init() {
	var err error
	// 打开与mysql的连接
	DB, err = sql.Open("mysql", g.Config().Database)
	if err != nil {
		log.Fatalln("open db fail:", err)
	}
    // 连接池设置,最大idle
	DB.SetMaxIdleConns(g.Config().MaxIdle)
	// 测试能否连接上
	err = DB.Ping()
	if err != nil {
		log.Fatalln("ping db fail:", err)
	}
}
