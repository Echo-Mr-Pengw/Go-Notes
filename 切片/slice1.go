// 切片的定义及初始化，切片是引用类型

package main 

import "fmt"

func main() {

	//第一种方法
	//定义并初始化数组
	arr := [3]string{"刘备", "关羽", "张飞"}
	//定义切片
	var s []string

	s = arr[0:2] //包含0 1 不包含2
	fmt.Println(s)  //[刘备 关羽]

	//第二种方法 使用make
	s1 := make([]int, 3)
	s1[0] = 1
	s1[1] = 2
	s1[2] = 3
	fmt.Println(s1)  // [1 2 3]

	//第三种方法
	s2 := []int{1, 2, 3}
	fmt.Println(s2)  // [1 2 3]
}