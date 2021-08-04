package channel_close

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

// 广播
// 生产者：创建一个协程，向 channel 放入 0-9，然后结束等待。
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// 数据生产完毕，通过关闭 channel 来通知消费者，降低耦合度
		close(ch)
		wg.Done()
	}()
}

// 消费者：创建一个协程，从 channel 中取出 10个数据，然后结束等待
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		id := rand.Intn(10)
		for {
			// 接受信息的同时，关注 channel 是否被关闭
			if data, ok := <-ch; ok {
				fmt.Println("dataReceiver ", id, " - ", data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

// 创建一个 channel，连接生产者和消费者，等待
// 往关闭的 channel 上发送信息，会 panic
// 接收关闭的 channel 的信息，收到 channel 的零值
func TestChannelClose(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	// 有了 channel close，我们可以使用多个 receiver
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()

}

// dataReceiver  1  -  0
// dataReceiver  1  -  2
// dataReceiver  7  -  1
// dataReceiver  1  -  3
// dataReceiver  1  -  5
// dataReceiver  1  -  6
// dataReceiver  1  -  7
// dataReceiver  1  -  8
// dataReceiver  1  -  9
// dataReceiver  7  -  4
