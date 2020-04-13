//数组作为参数传递（值传递和引用传递）

package main 

import "fmt"

// 值传递
func test(arr1 [3]int) {

	arr1[0] = 11
	fmt.Println(arr1)  // [11 2 3]
}

// 引用传递
func test1(arr1 *[3]int) {
	arr1[0] = 22
	fmt.Println(arr1)  // [22 2 3]
}

func main() {

	arr1 := [3]int{1, 2, 3}

	//值传递 即使test方法修改了数组的值，原数组的值不变
	test(arr1)
	fmt.Println(arr1)  // [1 2 3]

	//引用传递
	test1(&arr1)
	fmt.Println(arr1)  // [22 2 3]
}