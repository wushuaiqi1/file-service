package utils

import (
	"context"
	"file-service/common"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"log"
)

//同步发送消息

func SendSync(topic string, body []byte) {
	message := primitive.NewMessage(topic, body)
	log.Println("SendSync req,", message)
	res, err := (*common.Producer).SendSync(context.Background(), message)
	if err != nil {
		fmt.Printf("send message error: %s\n", err)
		return
	}
	fmt.Printf("send message success: result=%s\n", res.String())
}
