// 错误异常处理
// defer panic recover

package main 

import "fmt"


func test() {

	//匿名函数
	fun := func() {
		err := recover()  //捕获异常
		if err != nil {
			fmt.Println("错误提示：", err);
		}
	}

	defer fun();

	num1 := 1
	num2 := 0

	res := num1 / num2
	fmt.Println(res)
}

func main() {

	test();
}