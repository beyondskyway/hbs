package main

import (
	"flag"
	"fmt"
	"github.com/open-falcon/hbs/cache"
	"github.com/open-falcon/hbs/db"
	"github.com/open-falcon/hbs/g"
	"github.com/open-falcon/hbs/http"
	"github.com/open-falcon/hbs/rpc"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	// 配置文件解析
	g.ParseConfig(*cfg)
	// 数据库连接池初始化
	db.Init()
	// 查询数据库记录缓存到内存
	cache.Init()
    // 清理hbs内存中过期(修改agent名称后)的agent记录
	go cache.DeleteStaleAgents()
	// http server
	go http.Start()
	// rpc server
	go rpc.Start()
    // 信号接收和销毁资源
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println()
		db.DB.Close()
		os.Exit(0)
	}()

	select {}
}
