package main

import (
	// "SimpleDouyin/controller"
	"SimpleDouyin/module"
	"SimpleDouyin/repository/mysql"
	"SimpleDouyin/service"
	"fmt"
	"github.com/gin-gonic/gin"
	// "net/http"
)

func main() {

	// 配置mysql连接
	fmt.Println(" I am here")
	if err := mysql.Init(); err != nil {
		fmt.Printf("connect failed, err: %v\n", err)
		return
	}
	defer mysql.Close()

	// 初始化snowflake算法
	if err := module.SnowflakeInit("2023-08-07", 1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	go service.RunMessageServer()

	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}


