// 主线程结束，如果其它子线程还未结束，整个程序会停止。
// 可以用time包，增加线程的睡眠时间，但是睡眠时间的大小不好定。
// 也可以用sync包里面的WaitGroup

package main

import (
	"fmt"
	"sync"
)

var w = sync.WaitGroup{}

func main() {

	w.Add(1) //将内部计数器的值设置1,也可以理解为开启1个线程
	go Write() // 开启写线程
	w.Wait() // 等到开启的线程都结束，也就是内部计数器等于0，主线程再执行后面的程序
	fmt.Println("所有的子线程都执行完毕了")
}

func Write() {
	defer w.Done() // 线程结束后，内部计数器-1;开始w.add(1),将内部计数器的值设置为1
	for i := 0; i <= 10; i++ {
		fmt.Println("write ", i)
	}
}
