package string_test

import "testing"

func TestString(t *testing.T) {
	// 声明 string
	var s string
	t.Log("为初始化的 s:", s)

	s = "hello"
	t.Log("hello:", len(s))

	// string 是不可变的 byte slice
	// s[1] = 'c'

	// 存放中文：严
	s = "严"
	t.Log("严:", len(s)) // 3

	// 对比 unicode 和 utf8
	// 数据类型：rune，可以从 string 里取出 unicode
	s = "中"
	c := []rune(s)
	t.Logf("中 unicode: %x", c[0]) // 4e2d 编码
	t.Logf("中 utf8: %x", s)       // e4b8ad 存储

}

func TestStringToRune(t *testing.T) {
	s := "小飞机，纸飞机"
	for _, c := range s {
		// [1] 表示 %c %d (二进制) 都是从 c 来取值
		t.Logf("%[1]c - %[1]x", c)
	}
}
