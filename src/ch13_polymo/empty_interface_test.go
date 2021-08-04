package polymo_test

import (
	"fmt"
	"testing"
)

// 空接口 和 断言
func DoSth(p interface{}) {
	// if i, ok := p.(int); ok {
	// 	fmt.Println("Integer:", i)
	// }
	// if i, ok := p.(string); ok {
	// 	fmt.Println("string:", i)
	// }
	// fmt.Println("Unknow type")

	// swith 写法
	switch v := p.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("string:", v)
	default:
		fmt.Println("Unknow type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSth(10)
	DoSth("10")
}

// 使用更小的接口，很多接口只包含一个方法 Reader.Read(), Writer.Write()
// 较大的接口，可以用小接口组合 ReaderWriter {Reader, Writer}
// 只依赖于必要功能的最小接口。用 Reader 就不要用 ReaderWriter。
