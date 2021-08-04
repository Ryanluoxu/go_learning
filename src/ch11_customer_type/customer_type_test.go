package cust_type

import (
	"fmt"
	"testing"
	"time"
)

//
// func timeSpent(inner func(op int) int) func(op int) int {
// 	return func(n int) int {
// 		start := time.Now()
// 		ret := inner(n)
// 		fmt.Println("time spent: ", time.Since(start).Seconds())
// 		return ret
// 	}
// }

// 将特定的方法，定义成一个类型：
type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv {
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

func TestTimeSpent(t *testing.T) {
	slowFuncWithTimeSpent := timeSpent(slowFunc)
	slowFuncWithTimeSpent(99)
}
