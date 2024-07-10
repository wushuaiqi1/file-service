package main

import (
	"log"
	"notice-service/common"
	"notice-service/mq"
)

func main() {
	noticeServiceStart()
}

func noticeServiceStart() {
	err := common.InitDatabase()
	if err != nil {
		return
	}

	err = mq.InitRocketMq()
	if err != nil {
		return
	}
	log.Println("notice service start success")
}
