//全局匿名函数

package main 

import "fmt"

//全局匿名函数
var Func = func (num1, num2 int) int {
	return num1 - num2
}

func main() {

	res := Func(1, 2)
	fmt.Println(res);  //output: -1
}