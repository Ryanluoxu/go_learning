package package_test

import (
	"testing"

	cm "github.com/easierway/concurrent_map"
)

// cd /Users/luoxu/go_learning/src/ch16_package
// go get github.com/easierway/concurrent_map

func TestPackage(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}

// 同一环境下，不同项目只能使用相同版本的包
// 无法实现特定版本的管理
// 使用 vender 来解决，在 src/vender
// godep, glide, dep 依赖管理工具

// glide:
// glide init
// glide install : create vendor
