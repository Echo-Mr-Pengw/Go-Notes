// defer 延时机制 
// 程序执行到defer后，把defer后面的内容值拷贝到栈中
// 当此函数执行完毕后，从栈中取出defer后面的代码执行

package main 

import "fmt"

func test() {

	defer fmt.Println("test1")
	defer fmt.Println("test2")

	fmt.Println("test3");
}

func main() {

	test()

	fmt.Println("test4");
}

// output：
// test3
// test2
// test1
// test4

