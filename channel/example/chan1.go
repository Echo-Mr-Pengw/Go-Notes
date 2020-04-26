// 管道的声明及简单使用
// 管道默认是双向-可读可写

package main 

import (
	"fmt"
)

func main() {

	// 第一种方式，管道默认是双向（可读可写）
	var ch chan int
	ch = make(chan int, 1) // 管道可存1个整型数据

	// 向管道写数据
	ch<- 10

	// 关闭管道
	close(ch)

	// 从管道取数据
	num := <-ch

	fmt.Println(num)  // output: 10


	// 第二种方式
	ch1 := make(chan int, 2) // 管道可存2个整型数据

	// 向管道写数据
	ch1<- 100
	
	// 关闭管道
	close(ch1)

	// 从管道取数据
	num1 := <-ch1
	num2 := <-ch1

	fmt.Println(num1, num2)  // output: 100 0
}