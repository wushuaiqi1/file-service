package tests

import (
	"log"
	"sync"
	"testing"
	"time"
)

var localMap sync.Map

func TestMapLock(t *testing.T) {

	for i := 0; i < 100; i++ {
		go func(userId string) {
			//如果存在 执行
			val, ok := localMap.LoadOrStore(userId, 1)
			if ok {
				log.Println("已经被读取:", val)
			} else {
				log.Println("已经存储：", val)
			}
		}("user:1")
	}
	time.Sleep(time.Second * 2)
}

func TestSyncMapStoreAndLoad(t *testing.T) {
	//map.put
	localMap.Store("user:1", 1)
	//mao.get
	val, ok := localMap.Load(1)
	if ok {
		log.Println(val)
	} else {
		log.Println(ok)
	}
	val, _ = localMap.Load("user:1")
	log.Println(val)
}
