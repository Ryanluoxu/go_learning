package groutine_test

import (
	"fmt"
	"testing"
	"time"
)

// 协程
// 初始化的 stack 只有2k。
// KSE 是 M：N
// 一个系统线程 -> 一个 processer -> 一串 Go协程。

// 线程对应 KE，切换起来开销比较大

// 协程的使用
func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Millisecond * 50)
}
