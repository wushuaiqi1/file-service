package main

import (
	"file-service/common"
	"file-service/routers"
	"file-service/utils"
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

	utils.SendSync(common.TopicFileUploadNotice, []byte("sss"))

	//初始化路由配置
	engine := routers.InitRouters()
	err = engine.Run("127.0.0.1:9090")
	if err != nil {
		log.Println("Gin Web Serve Start Fail,", err)
		return
	}
}
