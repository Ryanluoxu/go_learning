package interface_test

import "testing"

// 只要方法签名一致即可
// 非入侵，实现不依赖于接口定义
// interface 可以打在 client 打包里

type Programmer interface {
	HelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) HelloWorld() string {
	return "fmt.Println(\"Hello world\")"
}

func TestInterface(t *testing.T) {
	var p Programmer = new(GoProgrammer)
	t.Log(p.HelloWorld())
}
