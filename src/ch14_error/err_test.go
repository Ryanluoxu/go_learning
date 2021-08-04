package err_test

import (
	"errors"
	"fmt"
	"testing"
)

var LessThanTwoError = errors.New("n should not be less than 2")
var LargerThanHundredError = errors.New("n should not be larger than 100")

// 及早失败，避免嵌套
func GetFibonacci(n int) ([]int, error) {

	// if n < 2 || n > 100 {
	// 	return nil, errors.New("n should be in [2,100]")
	// }

	// 区分错误类型
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThanHundredError
	}

	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	t.Log(GetFibonacci(10))

	// 两段式
	if v, err := GetFibonacci(-10); err != nil {
		if err == LessThanTwoError {
			fmt.Println("it is less than 2")
		}
	} else {
		t.Log(v)
	}
}
