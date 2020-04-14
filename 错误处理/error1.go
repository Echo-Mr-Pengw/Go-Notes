// 错误处理机制  defer + panic + recover
// 如果不捕获处理异常错误，产生错误后面的代码不能执行

package main 

import "fmt"
	
// 错误处理函数
func errorHandle() {

	deferFunc := func () {
		err := recover() // recover是内置函数 可以捕获异常
		if err != nil {
			fmt.Println("error occurred")
		}
	}

	defer deferFunc()

	num1 := 1
	num2 := 0

	res := num1 / num2

	fmt.Println(res)
}

func main() {

	errorHandle()  // 如果这里不捕获处理异常错误，后面的代码不执行
	fmt.Println("next.....")

}

// output:
// error occurred
// next.....