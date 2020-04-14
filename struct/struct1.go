//结构体的定义及声明

package main 

import "fmt"

type User struct {
	Name string
	Age int
	Sex string
}

func main() {

	// 第一种方法
	// 声明u变量类型为结构体
	var u User
	fmt.Println(u)  // { 0 }

	// 赋值
	u.Name = "new2014kb"
	u.Age  = 19
	u.Sex  = "男"

	// 输出全部的值
	fmt.Println(u) // {new2014kb 19 男}

	// 输出单个值
	fmt.Println(u.Name, u.Age, u.Sex)  // new2014kb 19 男

	// 第二种方法
	// 定义并初始化
	u1 := User{"张飞", 20, "男"}
	// 输出值
	fmt.Println(u1, u1.Name, u1.Age, u1.Sex)  // {张飞 20 男} 张飞 20 男
}