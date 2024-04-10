package main

import (
	"file-service/common"
	"file-service/routers"
	"log"
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

	//初始化MQ
	err = common.InitRocketMQ()
	if err != nil {
		log.Println("InitRocketMQ Error...", err)
		return
	}

	//初始化路由配置
	engine := routers.InitRouters()
	err = engine.Run()
	if err != nil {
		log.Println("Gin Web Serve Start Fail,", err)
		return
	}
	log.Println("Gin Web Serve Start Success...")
}
