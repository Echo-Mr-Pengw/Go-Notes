// 声明单向管道 可写或者可读

package main;

import (
	"fmt"
)

func main() {

	// 声明管道为只写
	var ch chan<- int
	ch = make(chan int, 1)

	// 向管道写入数据
	ch<- 1

	close(ch)

	// 从管道读数据
	// num := <-ch  //如果读取或报错  .\chan2.go:20:9: invalid operation: <-ch (receive from send-only type chan<- int)


	//声明管道为只读
	var ch1 <-chan int
	ch1 = make(chan int, 1)

	// 向管道写入数据
	// ch1<- 1  // .\chan2.go:29:5: invalid operation: ch1 <- 1 (send to receive-only type <-chan int)

	//fmt.Println(ch1)
}

