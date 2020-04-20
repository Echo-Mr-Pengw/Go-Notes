// 自定义类型实现多个接口

package main 

import(
	"fmt"
)

// A接口
type A interface {
	Say()
}

// B接口
type B interface {
	Play()
}

type C struct {

}

func (c C) Say() {
	fmt.Println("Say() ....")
}

func (c C) Play() {
	fmt.Println("Play() ....")
}

func main() {

	var c C

	var a A = c 
	var b B = c

	a.Say()
	b.Play()
}