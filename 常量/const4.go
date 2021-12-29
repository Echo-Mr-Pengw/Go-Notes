// 常量两种使用方式，对编译器的影响

package main

import "fmt"

func main() {

	// 未声明类型
	const a = 100
	const b float64 = a
	fmt.Printf("%T, %v, %T, %v", a,a,b,b)  // int, 100, float64, 100

	// 声明类型
	const c int = 100
	const d float64 = c
	fmt.Printf("%T, %v, %T, %v", c,c,d,d) // cannot use c (type int) as type float64 in const initializer
}