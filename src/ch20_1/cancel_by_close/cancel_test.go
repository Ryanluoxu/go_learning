package cancel_test

import (
	"fmt"
	"testing"
	"time"
)

// CSP 模型下的任务取消
func TestCancel(t *testing.T) {
	cancelChannel := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		// 创建协程，将 i 和 channel 传入
		go func(i int, cancelCh chan struct{}) {
			// 直到 channel cancel，才结束协程
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println("TestCancel - grt ", i, " - end..")
		}(i, cancelChannel)
	}
	// cancel_1(cancelChannel)	// 只有一个协程接收到信号，得以结束
	cancel_2(cancelChannel) // 所有协程都可以结束
	time.Sleep(time.Second * 1)

}

// 收到消息，表示被取消？
func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

// 发送任意东西，表示被取消？
// 只有一个协程接收到信号，得以结束
func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

// 可以取消所有的协程
// 自带的广播机制
func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}
