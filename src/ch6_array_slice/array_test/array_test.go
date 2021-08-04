package array_test

import "testing"

func TestArray(t *testing.T) {

	var arr [3]int
	t.Log(arr) // [0 0 0]

	arr1 := [...]int{1, 2, 3, 4}
	t.Log(arr1)
	arr1[2] = 5
	t.Log(arr1)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}

	// go 只有 for 来遍历。
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	// 类似 java 的 for each
	for idx, e := range arr3 {
		t.Log(idx, e)
	}

	// 对 index 不感兴趣。
	for _, e := range arr3 {
		t.Log(e)
	}

}

// 数组的截取：前闭后开
func TestArraySection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}

	// 3 - end：[4 5]
	arr_sec := arr[3:]
	t.Log(arr_sec)

	// start - 2： [1 2 3]
	arr_sec = arr[:3]
	t.Log(arr_sec)

	// 不支持 [:-1]
}
