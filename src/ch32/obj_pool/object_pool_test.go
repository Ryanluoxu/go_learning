package obj_pool

import (
	"fmt"
	"testing"
	"time"
)

func TestObjPool(t *testing.T) {

	pool := NewObjPool(10)
	t.Log("TestObjPool - pool: ", pool)

	// 尝试获得 11 个对象
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error("TestObjPool - error:", err)
		} else {
			fmt.Println("TestObjPool - get object: ", &v)

			// release
			// if err := pool.ReleaseObj(v); err != nil {
			// 	t.Error("TestObjPool - err:", err)
			// }
			// 不放回就会出现 timeout
		}
	}

	fmt.Println("TestObjPool - end..")
}
