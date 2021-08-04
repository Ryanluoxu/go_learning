package slice_test

import "testing"

/**
 * 切片：指针，元素个数，数组容量
 * 类似 java 的 arrayList
 */
func TestSliceInit(t *testing.T) {
	// 差别只在于[]没有声明长度
	var s0 []int
	t.Log("初始的切片（长度，容量）：", len(s0), cap(s0))

	s0 = append(s0, 1)
	t.Log("append 1 之后：", len(s0), cap(s0))

	// 另一张声明方式：make 类型，长度，容量
	// 对于使用者来说，只会关注 length，容量只在不足时扩充
	s2 := make([]int, 3, 5)
	t.Log("make([]int, 3, 5)：", len(s2), cap(s2))

	t.Log("s2[0]:", s2[0])
	t.Log("s2[1]:", s2[1])
	t.Log("s2[2]:", s2[2])

	s2 = append(s2, 1)

	t.Log("append(s2, 1): ", len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2], s2[3])
}

// 展示切片如何变长
// 容量不足，则 *2：1248
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		// 有可能是返回一个新的地址，所以要 s = append
		// 自增长的代价
		s = append(s, i)
		t.Log("append", i, ": ", len(s), cap(s))
	}
	/*
	  slice_test.go:23: 1 1
	  slice_test.go:23: 2 2
	  slice_test.go:23: 3 4
	  slice_test.go:23: 4 4
	  slice_test.go:23: 5 8
	  slice_test.go:23: 6 8
	  slice_test.go:23: 7 8
	  slice_test.go:23: 8 8
	  slice_test.go:23: 9 16
	  slice_test.go:23: 10 16
	*/
}

// 展示切片如何共享内存
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	q2 := year[3:6]
	t.Log("year[3:6]:", len(q2), cap(q2)) // 3 9
	// 由此可见，len 是针对当前的变量，但是 cap 对应的是存储的切片的容量。

	summer := year[5:8]
	t.Log("year[5:8]:", len(summer), cap(summer)) // 3 7

	summer[0] = "changed"
	t.Log("summer[0] = changed:")
	// 因为是共享内存，所以修改对其他变量产生影响：
	t.Log(q2)   // [Apr May changed]
	t.Log(year) // [Jan Feb Mar Apr May changed Jul Aug Sep Oct Nov Dec]
}

// 数组是可以比较的：前提是相同长度，相同容量，里面的元素都相同
// 切片：不可以
func TestSliceCompare(t *testing.T) {
	// slice can only be compared to nil
	//
	// a := []int{1, 2, 3, 4}
	// b := []int{1, 2, 3, 4}
	//
	// if a == b {
	// 	t.Log("equal")
	// }
}
