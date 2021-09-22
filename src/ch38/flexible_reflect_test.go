package reflect_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

type Customer struct {
	CookieId string
	Name     string
	Age      int
}

func TestDeepEqual(t *testing.T) {
	a := map[int]int{1: 1, 2: 2, 3: 4}
	b := map[int]int{1: 1, 2: 2, 3: 4}
	// fmt.Println("TestDeepEqual - a == b:", a == b)	// (map can only be compared to nil)
	fmt.Println("TestDeepEqual - a DeepEqual b:", reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}
	fmt.Println("TestDeepEqual - s1 DeepEqual s2:", reflect.DeepEqual(s1, s2))
	fmt.Println("TestDeepEqual - s1 DeepEqual s3:", reflect.DeepEqual(s1, s3)) // false

	c1 := Customer{"1", "Mike", 30}
	c2 := Customer{"1", "Mike", 30}
	fmt.Println("TestDeepEqual - c1 DeepEqual c2:", reflect.DeepEqual(c1, c2)) // true
	fmt.Println("TestDeepEqual - c1 == c2:", c1 == c2)                         // true
}

type Employee struct {
	EmployeeId string
	Name       string `format:"hahaha"` // annotation
	Age        int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

// 通用的方法来填充不同的 type
func fillBySetting(st interface{}, settings map[string]interface{}) error {

	// st must be a prt
	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		// Elem() 获取指针指向的结构
		if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
			return errors.New("fillBySetting - 1st param should be a pointer to the struct type")
		}
	}

	// 核心
	var (
		field reflect.StructField
		ok    bool
	)

	for k, v := range settings {
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}

	return nil
}

func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name": "Tom", "Age": 22}

	e := Employee{}
	if err := fillBySetting(&e, settings); err != nil {
		t.Fatal("TestFillNameAndAge - err:", err)
	}
	t.Log("TestFillNameAndAge - e:", e)

	c := new(Customer)
	if err := fillBySetting(c, settings); err != nil {
		t.Fatal("TestFillNameAndAge - err:", err)
	}
	t.Log("TestFillNameAndAge - c:", c)

}
