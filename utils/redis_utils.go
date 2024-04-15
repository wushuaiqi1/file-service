package utils

import (
	"context"
	"file-service/common"
	"log"
	"strings"
	"time"
)

const lockPrefix = "lock:"

func RedisLock(key string, value string, expireTime time.Duration) bool {
	if !strings.HasPrefix(key, lockPrefix) {
		log.Println("RedisLock key of string must lock as prefix,your key:", key)
		return false
	}
	ok, err := common.RedisInstance.SetNX(context.Background(), key, value, expireTime).Result()
	if err != nil {
		log.Println("RedisLock get lock fail:", err)
		return false
	}
	if ok {
		log.Println("RedisLock get lock success:", ok)
	}
	return ok
}

func RedisUnlock(key string, value string) {
	if !strings.HasPrefix(key, lockPrefix) {
		log.Println("RedisUnlock key of string must lock as prefix,your key:", key)
		return
	}
	val, err := common.RedisInstance.Get(context.Background(), key).Result()
	if err != nil || val != value {
		log.Println("RedisUnLock unlock fail,", err)
		return
	}
	res := common.RedisInstance.Del(context.Background(), key)
	if res.Val() > 0 {
		log.Println("RedisUnLock unlock success,rows:", res.Val())
	}
}
