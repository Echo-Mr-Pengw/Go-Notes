// 变量退化成赋值
// 相同的变量名，全局和局部变量不影响

package main

import "fmt"

// 定义全局变量
var g = 100

func main() {

	//演示： 相同的变量名，全局和局部变量不影响
	fmt.Println(&g, g)  // 0x11562a8 100

	g := 101
	fmt.Println(&g, g) // 0xc0000b4008 101

	// 演示：变量退化成赋值
	i := 10
	fmt.Println(&i, i)    // 0xc00001a0c0 10

	i, j := 11, 12
	fmt.Println(&i, i, j) // 0xc00001a0c0 11 12
}

