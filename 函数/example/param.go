// 变参，必须是切片且未同类型

package main

import "fmt"

func main() {

	// 第一种
	// test1("h", 1, 2, 3)

	// 切片
	//str := "h"
	//s := []int{1, 2, 3}
	//test1(str, s...)

	// 数组
	str := "h"
	arr := [2]int{1, 2}
	test1(str, arr[:]...)  // 数组转变成切片，然后展开

}

func test1(str string, s ...int) {
	fmt.Println(str, s)
}