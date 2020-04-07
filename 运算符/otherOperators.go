// *和&运算符

package main 

import "fmt"

func main() {

	i := 1

	var ptr *int = &i

	//& 取地址  *表示输出指针指向地址的值
	fmt.Println(&i, *ptr);
}