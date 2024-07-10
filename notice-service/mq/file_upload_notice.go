package mq

import (
	"context"
	"encoding/json"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"log"
	"notice-service/models"
	"notice-service/repository"
	"notice-service/utils"
	"time"
)

func InitRocketMq() (err error) {
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("test"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
	)
	if err != nil {
		return nil
	}

	if err := c.Subscribe(
		TopicFileUploadNotice,
		consumer.MessageSelector{},
		// 收到消息后的回调函数
		func(ctx context.Context, args ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for i := range args {
				log.Printf("Topic:FileUploadNotice获取到值： %v \n", args[i])
				body := args[i].Body
				notice := models.Notice{}
				err := json.Unmarshal(body, &notice)
				if err != nil {
					return 0, err
				}
				log.Printf("Topic:FileUploadNotice 解析body: %v \n", notice)
				//更新数据库
				nr := repository.NewNoticeRepository()
				nr.Create(&notice)
				//机器人消息通知
				err = utils.RobotSendTextMessage(string(args[i].Body))
				if err != nil {
					return 0, err
				}
			}
			return consumer.ConsumeSuccess, nil
		}); err != nil {
		log.Println("Subscribe Error:", err)
		return nil
	}
	err = c.Start()
	if err != nil {
		panic("启动consumer失败")
	}
	//不能让主goroutine退出
	time.Sleep(time.Hour)
	_ = c.Shutdown()
	return nil
}

const TopicFileUploadNotice string = "FileUploadNotice"
