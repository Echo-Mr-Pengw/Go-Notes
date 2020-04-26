// 数组元素的获取及遍历

package main 

import (
	"fmt"
)

func main() {

	arr := [3]int{1, 2, 3}

	// 第一种方法
	arrLen := len(arr)

	for i := 0; i < arrLen; i++ {
		fmt.Println(arr[i])
	}

	// 第二种方法
	// i 是key  v是值
	for i, v := range arr {
		fmt.Println(i, v)
	}
}