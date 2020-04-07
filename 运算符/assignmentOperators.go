//赋值运算符

package main 

import "fmt"

func main() {

	//声明并初始化变量
	i := 1

	// +=  
	// 等价于 i = i + 2
	i += 2
	fmt.Println(i);    //output: 3

	// -=
	// 等价于 i = i - 2
	i -= 2
	fmt.Println(i);    //output: 1

	// *=
	// 等价于 i = i * 2
	i *= 2
	fmt.Println(i);   //output: 2

	// /=
	// 等价于 i = i / 2
	i /= 2
	fmt.Println(i);   //output: 1

	// %=
	// 等价于 i = i % 2
	i %= 2
	fmt.Println(i);   //output: 1
}