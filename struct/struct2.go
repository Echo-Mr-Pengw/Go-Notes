// 结构体 值类型

package main 

import "fmt"

type User struct {
	Name string
	Age int
	Sex string
}

func main() {

	// 声明并初始化
	u1 := User{"张飞", 30, "男"}
	// 赋值
	u2 := u1

	fmt.Println(u1, u2) // {张飞 30 男} {张飞 30 男}

	//修改值
	u1.Name = "关羽"
	u1.Age  = 40

	fmt.Println(u1, u2) // {关羽 40 男} {张飞 30 男}
}