package share_mem

import (
	"sync"
	"testing"
	"time"
)

func TestShareMem(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		// 线程不安全
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter) // 4758
}

// 线程安全，互斥锁
func TestThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)     // 需要时间把协程执行完
	t.Logf("counter = %d", counter) // 5000
}

// waitGroup 只有当等待的东西都完成了，才能继续执行
func TestWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1) // 增加一个协程，就要多一个计数
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done() // 结束一个协程，通知 wg
		}()
	}
	// 用 wg 替换 sleep, 大大节约时间
	// time.Sleep(1 * time.Second)  // 需要时间把协程执行完
	wg.Wait()
	t.Logf("counter = %d", counter) // 5000
}
