// 继承使用细节之一
// 结构体可以使用嵌套匿名结构体中的所有字段和方法（不区分大小写）


package main 

import (
	"fmt"
)

type Stu struct {
	Name string  // 字符名大写
	age int      // 字符名小写
}
 
// 方法名大写
func (a Stu) Say() {
	fmt.Println("Say() ", a.Name, a.age)
}

// 方法名小写
func (a Stu) eat() {
	fmt.Println("eat() ")
}

type Puple struct {
	Stu   // 继承Stu结构体
}

func main() {

	var p Puple
	
	p.Name = "小明"
	p.age  = 20

	p.Say()
	p.eat()

	// output:
	// Say()  小明 20
	// eat()
}