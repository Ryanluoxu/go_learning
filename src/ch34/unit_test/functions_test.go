package testing

import (
	"fmt"
	"testing"

	// 表格测试法
	"github.com/stretchr/testify/assert"
)

func TestSquare(t *testing.T) {
	// 表格测试法
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		if ret != expected[i] {
			t.Errorf("TestSquare - input is %d, expected is %d, the actual is %d", inputs[i], expected[i], ret)
		}
	}
}

// Fail, Error: 测试继续
// Fatal: 该测试终止，其他测试继续
// 代码覆盖率： go test -v -cover
func TestError(t *testing.T) {
	fmt.Println("TestError - start")
	t.Error("TestError - error")
	t.Fatal("TestError - Fatal")
	fmt.Println("TestError - end") // error have, fatal no
}

// assert
// go get -u github.com/stretchr/testify
func TestSquareAssert(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		assert.Equal(t, expected[i], ret)
	}
}
