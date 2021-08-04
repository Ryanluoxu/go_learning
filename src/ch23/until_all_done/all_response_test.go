package all_response_test

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

// CSP 实现返回所有结果
func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	ret := ""
	for j := 0; j < numOfRunner; j++ {
		ret += <-ch + "\n"
	}
	return ret
}

func TestAllResponse(t *testing.T) {
	t.Log("TestAllResponse - before, NumGoroutine: ", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(time.Second * 1)
	t.Log("TestAllResponse - after, NumGoroutine: ", runtime.NumGoroutine())
	// TestFirstResponse - after, NumGoroutine:  11。剩下的协程阻塞。通过 buffer channel 来缓解
}
