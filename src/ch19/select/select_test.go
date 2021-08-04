package select_test

import (
	"fmt"
	"testing"
	"time"
)

// 多渠道， 超时机制
func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log("TestSelect - ", ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("TestSelect - timeout..")
	}
}

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "service completed.."
}

// 被调用的时候，是启用另一个协程，而不阻塞当前协程
// 当返回结果的时候，返回 channel，调用者需要结果就在 channel 里获取
func AsyncService() chan string {
	// buffer channel 避免因为调用者不在，自身被阻塞
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("AsyncService - get ret and return to channel..")
		retCh <- ret
		fmt.Println("AsyncService - end..")
	}()
	return retCh
}
