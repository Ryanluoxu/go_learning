package cancel_by_context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCancel(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 5; i++ {
		// 创建协程，将 i 和 ctx 传入
		go func(i int, ctx context.Context) {
			// 直到 ctx cancel，才结束协程
			for {
				if isCancelled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println("TestCancel - grt ", i, " - end..")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)

}

// 收到消息，表示被取消？
func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
