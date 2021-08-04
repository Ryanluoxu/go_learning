package object_test

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestNewObject(t *testing.T) {
	// 1.
	e1 := Employee{"1", "Ryan", 30}

	// 2.
	e2 := Employee{Id: "2", Name: "John", Age: 40}

	// 3.
	e3 := new(Employee) // 返回指针
	e3.Id = "3"
	e3.Name = "Wong"
	e3.Age = 35

	t.Log("e1:", e1)
	t.Log("e2:", e2)
	t.Log("e3:", e3)
	t.Logf("e2 is %T", e2) // struct
	t.Logf("e3 is %T", e3) // pointer, same as &e2
}

type Employee struct {
	Id   string
	Name string
	Age  int
}

// 对于行为的添加和定义: 两种方法有同样的效果，但是不能共存。
// 区别在于
// 1. 指针方法，不会对 emp 进行复制，方法内部的 emp 与外部调用的 emp 指向同一个地址
// 2. 结构方法，会对 emp 进行复制，方法内部的 emp 与外部调用的 emp 指向不同地址

func (emp *Employee) String() string {
	fmt.Printf("address is: %x", unsafe.Pointer(&emp.Name))
	return fmt.Sprintf("Employee{Id:%s, Name:%s, Age:%d}", emp.Id, emp.Name, emp.Age)
}

// func (emp Employee) String() string {
// 	return fmt.Sprintf("Employee{Id:%s, Name:%s, Age:%d}", emp.Id, emp.Name, emp.Age)
// }

func TestStructOperations(t *testing.T) {
	e := Employee{Id: "1", Name: "Ryan", Age: 30}
	t.Log(e.String())
	e1 := &e // 指针也是一样的
	t.Log(e1.String())
}
