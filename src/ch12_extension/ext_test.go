package ext_test

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (pet *Pet) Speak() {
	fmt.Print("...")
}

func (pet *Pet) SpeakTo(someone string) {
	pet.Speak()
	fmt.Println(" ", someone)
}

// 匿名嵌套类型
type Dog struct {
	// 感觉就像 extends Pet
	Pet
}

func (dog *Dog) Speak() {
	fmt.Print("wang..")
}

// 测试
func TestDog(t *testing.T) {
	// var dog Pet = new(Dog) // cannot use new(Dog) (type *Dog) as type Pet
	// 不支持 LSP
	dog := new(Dog)
	// dog := new(Dog)
	dog.SpeakTo("John")
}
