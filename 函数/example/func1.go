// 函数的定义及使用
// 函数名和变量名严格区分大小写，首字母大小表示公有的；反之私有的
// 函数可返回多个值

package main

import "fmt"

//定义函数，求两个数的和差
func sumAndSub(num1 int, num2 int) (int, int) {
	
	sum := num1 + num2
	sub := num1 - num2

	return sum, sub

}

func main() {

	sum, sub := sumAndSub(1, 2);
	fmt.Println(sum, sub);   //output： 3 -1
}