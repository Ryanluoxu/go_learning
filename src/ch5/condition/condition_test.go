package condition_test

import "testing"

func TestCondition(t *testing.T) {
	// if v, err := someFun(); err == nil {
	// 	t.Log(v)
	// } else {
	// 	t.Log("err")
	// }
}

func TestSwitch(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even: ")
		case 1, 3:
			t.Log("Odd:")
		default:
			t.Log("not 0-3")
		}
	}
}

func TestSwitch2(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even: ")
		case i%1 == 0:
			t.Log("Odd:")
		default:
			t.Log("...")
		}
	}
}
