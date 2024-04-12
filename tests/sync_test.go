package tests

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var id int
var mutex sync.Mutex

func TestLock(t *testing.T) {
	for i := 0; i < 1000; i++ {
		go func() {
			mutex.Lock()
			id++
			mutex.Unlock()
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Lock ID Res->", id)
}

func TestNoLock(t *testing.T) {
	for i := 0; i < 1000; i++ {
		go func() {
			id++
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("No Lock ID Res->", id)
}
