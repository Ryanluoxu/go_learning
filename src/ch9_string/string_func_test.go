package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFunc(t *testing.T) {
	// string 分割
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for idx, part := range parts {
		t.Log(idx, "-", part)
	}

	// 连接
	t.Log("join: ", strings.Join(parts, "+"))
}

// string 的转换
func TestStringConv(t *testing.T) {
	// 整数到 stirng
	s := strconv.Itoa(10)
	t.Log("str:" + s)

	// string 转整数
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}

}
