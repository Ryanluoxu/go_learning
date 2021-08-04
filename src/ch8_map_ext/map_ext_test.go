package map_ext

import "testing"

// 在 go 里，方法是一等公民
// map 里的 value 可以是方法。实现工厂模式。
func TestMapWithFuncValue(t *testing.T) {

	// 方法： func(variable int) int
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }

	t.Log("m[1](3):", m[1](3))
	t.Log("m[2](3):", m[2](3))
	t.Log("m[3](3):", m[3](3))
}

// 实现 set，map[type]bool
func TestMapForSet(t *testing.T) {
	set := map[int]bool{}
	set[1] = true

	// 判断一个元素是否在 set 里
	n := 2
	if set[n] {
		t.Log(n, "is existing.")
	} else {
		t.Log(n, "is not existing.")
	}

	// 获取元素的个数
	set[2] = true
	t.Log("len(set):", len(set))

	// 删除某个元素
	delete(set, 1)
	n = 2
	if set[n] {
		t.Log(n, "is existing.")
	} else {
		t.Log(n, "is not existing.")
	}

}
