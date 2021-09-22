package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

// 使用 type / value
func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log("TestTypeAndValue - type, value = ", reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log("TestTypeAndValue - type = ", reflect.ValueOf(f).Type())
}

// 使用 kind
func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("CheckType - float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("CheckType - Int")
	default:
		fmt.Println("CheckType - Others")
	}
}

func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(f)  // float
	CheckType(&f) // Others
}

// 利用反射编写灵活代码
type Employee struct {
	EmployeeId string
	Name       string `format:"hahaha"` // annotation
	Age        int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

type Customer struct {
	CookieId string
	Name     string
	Age      int
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}

	// 按名字获取成员
	name := reflect.ValueOf(*e).FieldByName("Name")
	t.Log("TestInvokeByName - name:", name)

	// struct Tag: 类似于 annotation
	if name, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("TestInvokeByName - fail")
	} else {
		t.Log("TestInvokeByName - Tag format:", name.Tag.Get("format"))
	}

	// 调用方法更新 Age
	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(10)})
	t.Log("TestInvokeByName - after upate:", e)
}
