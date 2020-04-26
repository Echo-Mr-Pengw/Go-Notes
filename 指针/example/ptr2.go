// 指针作为参数传递

package main

import (
	"fmt"
)

// 修改局部变量i的方法
func ptrParams(ptr *int) int {
	*ptr = *ptr + 10
	return *ptr
}

func main() {

	// 声明变量
	i := 10

	// 指针作为参数传递
	res := ptrParams(&i);
	fmt.Println(i, res);   //output: 20  20
}