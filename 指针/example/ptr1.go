// 指针的定义及变量地址
 package main 

import (
	"fmt"
)

func main() {

	// 定义变量
	i := 12
	// 定义指针 默认值为nil
	var ptr *int
	
	// &i 表示i的内存地址，指针指向的是地址
	ptr = &i
	// *ptr表示指针指向地址的内容，也就是i的值
	ptr_addr := *ptr

	fmt.Println(ptr, ptr_addr);   //prt是i的内容地址，ptr_add是i的值


}