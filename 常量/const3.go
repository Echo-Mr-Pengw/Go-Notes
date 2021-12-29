// 常量与变量的区别
// 常量只可读，且不能获取其地址

package main

import "fmt"

func main() {

	var i int
	const j = iota
	fmt.Println(&i, &j) // cannot take the address of a
}