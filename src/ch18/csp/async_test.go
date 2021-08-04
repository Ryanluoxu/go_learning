package async_test

import (
	"fmt"
	"testing"
	"time"
)

// CSP 并发机制
// - 通过 channel 通讯，耦合更松
// - 1.channel 两端的人需要同时在，进行通信。
// - 2.buffer channel。

func service() string {
	time.Sleep(time.Microsecond * 50)
	return "service done.."
}

func otherTask() {
	fmt.Println("working on something else..")
	time.Sleep(time.Microsecond * 100)
	fmt.Println("other task is done..")
}

// 串行
func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

// 被调用的时候，是启用另一个协程，而不阻塞当前协程
// 当返回结果的时候，返回 channel，调用者需要结果就在 channel 里获取
func AsyncService() chan string {
	// buffer channel 避免因为调用者不在，自身被阻塞
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("get ret and return to channel..")
		retCh <- ret
		fmt.Println("AsyncService end..")
	}()
	return retCh
}

// 并行，类似于 future
func TestAsync(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
	time.Sleep(time.Microsecond * 1000)
}
