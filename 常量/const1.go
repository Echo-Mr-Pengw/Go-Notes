// 常量数组中如果为指定类型和值，其类型跟上面类型相同

package main

import "fmt"

func main() {

	const (
		a = 1     // 第一个变量 不能即未指定变量类型又没赋值
		b
		c = "a"
		d
	)
	fmt.Printf("%T, %T, %T, %T", a, b, c, d) // int, int, string, string%
}
