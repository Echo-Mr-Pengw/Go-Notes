// 匿名结构体字段
// 注意：匿名结构体类型不能重复
package main

import "fmt"

type Person struct {
	string // 姓名
	int    // 年龄
	// 如果在定义一个string，则会报错
}

func main() {
	p := Person{"pengw", 20}
	fmt.Println(p, p.string, p.int)  // {pengw 20} pengw 20
}

