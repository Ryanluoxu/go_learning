package once_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

var singleInstance *Singleton
var once sync.Once

func GetSingletonObj() *Singleton {
	// 只运行一次
	once.Do(func() {
		fmt.Println("GetSingletonObj - new(Singleton)..")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestSingleton(t *testing.T) {
	// 在多个协程里调用
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Println("TestSingleton - ", i, unsafe.Pointer(obj)) // 相同地址
			wg.Done()
		}()
		wg.Wait()
	}
}
