//指向指针的指针

package main 

import "fmt"

func main() {

	//声明并初始化变量
	i := 1

	//指针
	var ptr1 *int

	//指向指针的指针
	var ptr2 **int

	//ptr1指向i的内存地址
	ptr1 = &i 

	//ptr2 指向ptr1指向的地址
	ptr2 = &ptr1

	//**ptr2 表示获取指向指针的指针地址，也就是i的地址
	fmt.Println(*ptr1, **ptr2);  //output: 1 1
}