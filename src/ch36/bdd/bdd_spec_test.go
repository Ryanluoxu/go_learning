package bdd

// BDD: behavior driven development
// story card: 背面写验收场景
// framework: goconvey
// 安装：go get -u github.com/smartystreets/goconvey/convey
// go get -u github.com/smartystreets/goconvey
// 启动 Web UI： $GOPATH/bin/goconvey

import (
	"testing"

	// 如果是 . 表示 import 到当前目录下，不需要 cv.xxx
	// . "github.com/smartystreets/goconvey/convey"
	cv "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {

	// only pass t to top-level Convey calls
	cv.Convey("Given 2 even numbers", t, func() {
		a := 3
		b := 4

		cv.Convey("When add the two numbers", func() {
			c := a + b

			cv.Convey("Then the result is still even", func() {
				cv.So(c%2, cv.ShouldEqual, 0)
			})
		})

	})
}
