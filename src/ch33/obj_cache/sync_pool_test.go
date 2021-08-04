package obj_cache

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

// sync pool，获取/放置对象的顺序：
// - 私有对象，协程安全
// - 共享池，协程不安全，使用时需要上锁
// - 如果前面两个都是空的，就会去别的 processer 的共享池里获取对象
// GC 会清空 pool

// 协程安全，会有锁的开销
// 生命周期受 GC 的影响，不适合做连接池等需要管理生命周期的资源
func TestSyncPool(t *testing.T) {

	// 创建一个 sync pool
	// 每次 pool 空的时候，都会 new 一个新的 100
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("TestSyncPool - create a new obj..")
			return 100
		},
	}

	// 获取 obj
	v := pool.Get().(int)
	fmt.Println("TestSyncPool - v:", v)

	// 往 pool 里放入一个 obj
	pool.Put(3)
	v1, _ := pool.Get().(int)
	fmt.Println("TestSyncPool - v1:", v1)

	// GC 清空 pool
	pool.Put(30)
	runtime.GC()
	v2, _ := pool.Get().(int) // 由于是空的 pool，所以会 create a new obj.
	fmt.Println("TestSyncPool - v2:", v2)
}

func TestSyncPoolInMultiGroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("TestSyncPoolInMultiGroutine - create a new obj..")
			return 100
		},
	}

	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup

	// 先取完了放进去的 3个 100， 然后 new obj
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			t.Log("TestSyncPoolInMultiGroutine - get pool obj:", pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
