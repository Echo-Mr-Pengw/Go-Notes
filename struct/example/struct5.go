// 匿名结构体
// 匿名结构体定义语法
// 变量 := struct{
// 		字段
// }{
// 		字段值
// }
package main

import(
	"fmt"
)

func main() {

	s1 := struct{
		Name string
		Age int
	}{
		Name : "pengw",
		Age: 20,
	}

	fmt.Println(s1, s1.Name, s1.Age) // {pengw 20} pengw 20
}