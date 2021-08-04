package polymo_test

import (
	"fmt"
	"testing"
)

type Code string
type Programmer interface {
	HelloWorld() Code
}

type GoProgrammer struct{}

func (p *GoProgrammer) HelloWorld() Code {
	return "fmt.Println(\"Hello World!\")"
}

type JavaProgrammer struct{}

func (p *JavaProgrammer) HelloWorld() Code {
	return "System.out.Println(\"Hello World!\")"
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.HelloWorld())
}

func TestPolymorphism(t *testing.T) {

	// interface 只能对应指针类型的实例
	goProg := new(GoProgrammer) // &GoProgrammer{}
	javaProg := new(JavaProgrammer)
	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)
}
