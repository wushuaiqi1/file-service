package main

import (
	"file-service/common"
	"file-service/routers"
	"log"
	"os"
)

func main() {
	webServiceStart()
}

func webServiceStart() {
	//初始化数据库
	err := common.InitDatabase()
	if err != nil {
		log.Println("InitDatabase Error...", err)
		return
	}
	err = common.InitRedis()
	if err != nil {
		log.Println("Redis Service Not Health:", err)
		return
	}

	//初始化MQ
	err = common.InitRocketMQ()
	if err != nil {
		log.Println("InitRocketMQ Error...", err)
		return
	}

	args := os.Args
	//初始化路由配置
	engine := routers.InitRouters()
	var port string
	if len(args) == 1 {
		port = "8080"
	} else {
		port = args[1]
	}
	err = engine.Run("127.0.0.1:" + port)
	if err != nil {
		log.Println("Gin Web Serve Start Fail,", err)
		return
	}
}
