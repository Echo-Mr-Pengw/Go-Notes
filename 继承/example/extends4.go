// 继承使用细节之三
// 结构体和其继承的嵌套匿名结构体出现字段和方法重复，则使用“就近原则”

package main 

import (
	"fmt"
)

// Stu结构体
type Stu struct {
	Name string
}

// Stu结构体中的方法
func (s Stu) Say() {
	fmt.Println("Stu结构体中Say() ", s.Name)
}

// Puple结构体
type Puple struct {
	Stu         //继承匿名结构体
	Name string
}

// Puple结构体中的方法
func (p Puple) Say() {
	fmt.Println("Puple结构体中的say()  ", p.Name)
}


func main() {

	var p Puple 

	p.Name = "张三"
	p.Say()  // output:  Puple结构体中的say()   张三
	

	// 如果需要使用Stu中的Name和Say()，使用匿名结构体的名字来区分调用
	p.Stu.Name = "李四"
	p.Stu.Say()   // output: Stu结构体中Say()  李四
}