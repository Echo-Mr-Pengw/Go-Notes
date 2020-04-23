// 匿名函数：没有函数名的函数

package main 

import "fmt"

func main(){

	i := func (num1, num2 int) int {
		return num1 - num2
	}

	fmt.Println(i(1, 2));  //output: -1
}