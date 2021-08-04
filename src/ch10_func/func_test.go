package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 返回多个数值
func TestReturnMultiValue(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
}

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

// 返回函数运行的时间, 通用的模式
// 输入是函数类型，返回也是函数类型
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent: ", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

// 1. 先创建一个函数
// 2. 再用 timeSpent 函数对这个函数进行加工，插入时间计算，返回一个新的函数
// 3. 用新的函数来运行，这样就可以得到运行的时长

// 《计算机程序的构造和解释》

func TestTimeSpent(t *testing.T) {
	slowFuncWithTimeSpent := timeSpent(slowFunc)
	slowFuncWithTimeSpent(99)
}

// 可变长参数
func sum(ops ...int) int {
	ret := 0
	for _, v := range ops {
		ret += v
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(sum(1, 2, 3))
	t.Log(sum(1, 2, 3, 4))
}

// defer 延迟执行函数
// 类似 try 的 finally，在函数结束前执行 defer
// panic：程序异常中断
func Clear() {
	fmt.Println("clear resources..")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start..")
	panic("error")
	// fmt.Println("after panic..") // 提示不会执行到
}
