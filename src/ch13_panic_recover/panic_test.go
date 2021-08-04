package panic

import (
	"errors"
	"fmt"
	"testing"
)

// panic and recover
func TestPanic(t *testing.T) {
	// recover
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover from: ", err)
		}
	}()
	fmt.Println("start..")
	panic(errors.New("panic.."))
}
