// 空接口。可以接收任意的数据类型

package main

import (
	"fmt"
)

type A interface {}

func main() {

	// 接口A 接收结构体
	var a1 A = struct {
		Name string
	}{
		Name: "pengw",
	}

	// 接口A 接收整型
	var a2 A = 12

	// 接口A 接收字符串
	var a3 A = "pengw"

	fmt.Println(a1, a2, a3) // {pengw} 12 pengw
}