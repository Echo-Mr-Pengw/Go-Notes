// 数组的初始化

package main 

import (
	"fmt"
)

func main() {

	// 第一种方法
	// 定义数组
	var arr1 [3]int
	//初始化数组
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3

	fmt.Println(arr1)

	// 第二种方法
	var arr2 [3]int = [3]int{11, 22, 33}
	fmt.Println(arr2)

	// 第三种方式
	var arr3 = [3]int{111, 222, 333}
	fmt.Println(arr3)

	// 第四种方式 ...表示不限定数组内值的个数
	var arr4 = [...]int{1111, 2222, 3333, 4444}
	fmt.Println(arr4)

	// 第五种方式
	arr5 := [3]int{11111, 22222, 33333}
	fmt.Println(arr5)

	// 第六种方式
	arr6 := [...]int{111111, 222222, 333333}
	fmt.Println(arr6)
}