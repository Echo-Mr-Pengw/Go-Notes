package main 

import (
	"fmt"
	"time"
)

func test() {

	for i := 0; i <= 10; i++ {
		fmt.Println("test() ", i)
	}

}

func main() {

	// 开启一个协程
	go test()

	// 等待1秒，不然看不到test函数打印的数据，
	// 因为，主线程执行完毕了，协程还在执行；主线程退出，协程也必须退出
	time.Sleep(time.Second)  

	for i := 0; i <= 10; i++ {
		fmt.Println("main() ", i)
	}
}