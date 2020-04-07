// switch 条件语句
// case后面可以跟多个表达式且用逗号隔开
// case后面可以不加break，默认系统会自动加

package main 

import "fmt"

func main() {

	//声明并初始化变量
	i := 1

	switch i {
		case 0:
			fmt.Println(0);
		case 1:
			fmt.Println(1);   //output: 1
		default:
			fmt.Println("no~~~~");
	}
}