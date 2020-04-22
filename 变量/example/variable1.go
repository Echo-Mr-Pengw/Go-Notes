// 定义单个变量的三种方式

package main

import (
	"fmt"
)

func main() {
	
	// 声明变量的第一种方式
	// 声明变量
	var i int

	// 变量赋值
	i = 100
	
	// 使用变量
	fmt.Println(i);

	// 明变量的第二种方式
	// 声明变量并赋值，由系统根据赋的值判断类型
	var j = 100

	// 使用
	fmt.Println(j);

	
	// 声明变量的第三种方式
	// 等同于 var k int   k = 100
	k := 100

	// 使用变量
	fmt.Println(k);
}


