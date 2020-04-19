// 继承使用细节之六
// 结构体如果嵌套一个有名结构体，再访问有名结构体内的方法和字段时，必须加上名字

package main 

import (
	"fmt"
)

type A  struct {
	Name string
}

type B struct {
	a A  // 嵌套一个有名结构体，
}

func main() {

	var b B

	// 这样写会报错，因为B嵌套的是一个有名结构体，
	// 必须通过名字来方法，b.a.Name = "张三"，否则报错
	// 如果B结构体中含有字段Name，则使用B中的Name
	// b.Name = "张三"
	
	b.a.Name = "张三"
	fmt.Println(b.a.Name)   // output: 张三
}
