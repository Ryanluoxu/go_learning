package unsafe_test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

func TestUnsafe(t *testing.T) {
	i := 10
	f := *(*float64)(unsafe.Pointer(&i))
	t.Log("TestUnsafe - ", unsafe.Pointer(&i))
	t.Log("TestUnsafe - ", f) // 5e-323
}

type Customer struct {
	Name string
	Age  int
}

// 合理的类型转换
type MyInt int

func TestConvert(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := *(*[]MyInt)(unsafe.Pointer(&a))
	t.Log("TestConvert - ", b)
}

// 原子类型操作
// 写在另一块地址，写好了之后，利用原子操作将 buffer 指向这块地址
func TestAtomic(t *testing.T) {
	// 初始化一个共享的 buffer 指针
	var shareBufPtr unsafe.Pointer

	// 写数据：向 data 写入0-99，然后将 buffer 指针指向 data
	writeDataFn := func() {
		data := []int{}
		for i := 0; i < 100; i++ {
			data = append(data, i)
		}
		atomic.StorePointer(&shareBufPtr, unsafe.Pointer(&data))
	}

	// 读数据：从 buffer 指针里读取数据
	readDataFn := func() {
		data := atomic.LoadPointer(&shareBufPtr)
		fmt.Println("TestAtomic - ", data, *(*[]int)(data))
	}

	var wg sync.WaitGroup
	writeDataFn()

	// 启动 10 个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			// 每个协程里，进行 10 次写数据操作
			for i := 0; i < 10; i++ {
				writeDataFn()
				time.Sleep(time.Millisecond * 100)
			}
			wg.Done()
		}()

	}
}
