package tests

import (
	"file-service/common"
	"file-service/utils"
	"log"
	"testing"
	"time"
)

func TestRedisLock(t *testing.T) {
	err := common.InitRedis()
	if err != nil {
		return
	}
	ok := utils.RedisLock("lock:", "1", time.Second*10)
	if ok {
		log.Println("获取分布式锁成功")
	} else {
		log.Println("获取分布式锁失败")
	}
	utils.RedisUnlock("lock:", "1")
}
