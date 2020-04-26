// 继承使用细节之四
// 结构体继承多个匿名结构体，且多个匿名结构体中含有相同的字段和方法;
// 在访问匿名结构体中的字段和方法时，需要加匿名结构体名称来访问


package main 

import (
	"fmt"
)

type A struct {
	Name string
}

type B struct {
	Name string
}

type C struct {
	A
	B
}

func main() {

	var c C

	// 如果这样访问则报错，因为C结构体本身没有Name字段，虽然A和B结构体有，但是编译器无法区分
	// 如果C结构体本身含有Name字段，则不会报错。编译器在C结构体就找到了，不会在去A和B中找
	//c.Name = "张三"
	
	//正确的方式
	c.A.Name = "张三"
	fmt.Println(c.A.Name) // output: 张三
}