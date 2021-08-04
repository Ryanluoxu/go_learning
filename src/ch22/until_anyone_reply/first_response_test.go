package first_response_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("return for %d", id)
}

// 同时跑 10 个协程，得到结果即刻放入 channel
// 一旦有数据，就会返回数据
// 没有 buffer 的 channel，容易阻塞
func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Log("TestFirstResponse - before, NumGoroutine: ", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second * 1)
	t.Log("TestFirstResponse - after, NumGoroutine: ", runtime.NumGoroutine())
	// TestFirstResponse - after, NumGoroutine:  11。剩下的协程阻塞。通过 buffer channel 来缓解
}
