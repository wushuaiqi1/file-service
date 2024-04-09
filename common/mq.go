package common

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

var Producer *rocketmq.Producer

func InitRocketMQ() error {
	p, err := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		producer.WithRetry(2), //指定重试次数
	)
	if err != nil {
		return err
	}
	if err = p.Start(); err != nil {
		return err
	}
	Producer = &p
	return nil
}
