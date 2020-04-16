//结构体的定义及声明

package main 

import "fmt"

type User struct {
	Name string
	Age int
	Sex string
}

func main() {

	// 第一种方式
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

	// 第二种方式
	// 定义并初始化
	u1 := User{"张飞", 20, "男"}
	// 输出值
	fmt.Println(u1, u1.Name, u1.Age, u1.Sex)  // {张飞 20 男} 张飞 20 男

	// 第三种方式
	// new后返回的是结构体指针，
	u2 := new(User)
	(*u2).Name = "关羽"  // 等价于 u2.Name = "关羽"
	(*u2).Age  = 20      // 等价于 u2.Age  = 20
	(*u2).Sex  = "男"    // 等价于 u2.Sex  = "男" 
	fmt.Println(u2)      // &{关羽 20 男}

	//第四种方式
	u3 := &User{"武松", 30, "男"}
	fmt.Println(u3)

	/**
	 * 第四种方式的另外一种写法
	 * u3 := &User{}
	 * (*u3).Name = "宋江"  // 等价于 u3.Name = "宋江"
	 * (*u3).Age  = 30      // 等价于 u3.Age  = 30
	 * (*u3).Sex  = "男"    // 等价于 u3.Sex  = "男" 
	 */

}