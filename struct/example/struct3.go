// 结构体 方法

package main 

import (
	"fmt"
)

type User struct {
	Name string
	Age int
	Sex string
}

func (u User) say() {
	fmt.Println(u.Name, u.Age, u.Sex)  //  张三 12 男
}

func main() {
	
	u := User{"张三", 12, "男"}
	//调用方法
	u.say()
}