// 继承使用细节之六
// 初始化时，给匿名结构体直接赋值

package main 

import (
	"fmt"
)


type A struct {
	Name string
	Age int
}

type B struct {
	A
}

func main() {

	b := B { 
		A {
			Name : "张",
			Age  : 10,
		},
	}

	fmt.Println(b.A) // {张 10}
}