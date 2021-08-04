package my_map

import "testing"

func TestMapInit(t *testing.T) {

	// 声明+赋值
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1)
	t.Log("len(m1): ", len(m1))

	// 先声明，再赋值
	m2 := map[int]int{}
	t.Log("m2:=map[int]int:", m2)

	m2[4] = 16
	t.Log("after m2[4]=16:", m2)

	// 用 make，可以一次性给足容量，减少之后的开销
	// 不可以用 cap 来查询 map
	m3 := make(map[int]int, 10)
	t.Log("m3:=make(map[int]int, 10)", m3)

}

func TestKeyNotExisting(t *testing.T) {

	// 访问不存在的 key
	m1 := map[int]int{}
	t.Log("visit non-exist key, value: ", m1[1])

	// 无法判断此 key 是否存在，所以：
	m1[3] = 111
	if v, ok := m1[3]; ok {
		t.Log("key 3 value:", v)
	} else {
		t.Log("key 3 is not existing")
	}

}

// 对 map 遍历
func TestMapTravel(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, " - ", v)
	}
}
