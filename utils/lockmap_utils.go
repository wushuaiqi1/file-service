package utils

import (
	"log"
	"sync"
	"time"
)

var localMap sync.Map

const val = "lockVal"

// GetLock 获取锁 获取成功返回true
func GetLock(key string) bool {
	_, ok := localMap.LoadOrStore(key, val)
	return !ok
}

// GetLockAndExpire 获取锁并设置锁失效时间
func GetLockAndExpire(key string, expireTime time.Duration) bool {
	getLockRes := GetLock(key)
	if !getLockRes {
		log.Println("获取锁失败，key:", key)
		return false
	}
	log.Println("获取锁成功，key:", key)
	go func() {
		time.Sleep(expireTime)
		UnLock(key)
	}()
	return true
}

// UnLock 释放锁
func UnLock(key string) {
	localMap.LoadAndDelete(key)
	log.Println("释放锁，key:", key)
}
