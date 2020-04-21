//一次性生成多个变量

package main

import "fmt"

func main() {

	//第一种方式
	//一次声明多个变量
	var i, j, k int
	//变量赋值
	i = 1
	j = 2
	k = 3
	//使用变量
	fmt.Println(i, j, k);

	//第二种方式
	var i1, j1, k1 = 11, 22, 33
	fmt.Println(i1, j1, k1);

	//第三中方法
	i2, j2, k2 := 111, 222, 333
	fmt.Println(i2, j2, k2);
}