// 继承的定义

package main 

import (
	"fmt"
)

// 声明学生结构体
type Student struct {
	Name string
	Age int
}

// 给学生结构体绑定方法
func (s Student) Say(name string , age int) {
	fmt.Println(name , age)
}

// 小学生
type PupilStudent struct {
	Student // 继承Student
}

// 高中生
type HightStudent struct {
	Student // 继承Student
}


func main() {

	// 小学生
	var p PupilStudent
	p_name := "张三"
	p_age := 10
	p.Student.Name = p_name   // 等价于 p.Name = p_name
	p.Student.Age  = p_age    // 等价于 p.Age = p_age

	p.Student.Say(p_name, p_age)

	// 高中生
	var h HightStudent
	h_name := "李四"
	h_age := 18
	h.Student.Name = h_name    // 等价于 h.Name = h_name
	h.Student.Age  = h_age     // 等价于 h.Age = h_age

	h.Student.Say(h_name, h_age)
}