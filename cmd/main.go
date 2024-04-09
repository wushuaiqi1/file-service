package main

import (
	"context"
	"file-service/common"
	"file-service/routers"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"log"
	"time"
)

func main() {

	consumerTest()
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

func consumerTest() {
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("test"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
	)
	if err != nil {
		panic(err)
	}
	if err := c.Subscribe("hello",
		consumer.MessageSelector{},
		// 收到消息后的回调函数
		func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for i := range msgs {
				fmt.Printf("获取到值： %v \n", msgs[i])
			}
			return consumer.ConsumeSuccess, nil
		}); err != nil {
	}
	err = c.Start()
	if err != nil {
		panic("启动consumer失败")
	}
	//不能让主goroutine退出
	time.Sleep(time.Hour)
	_ = c.Shutdown()
}
