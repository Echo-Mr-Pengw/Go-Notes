//goto 跳转到需要执行的语句

package main 

import (
	"fmt"
)

func main() {

	i := 1

	if i > 0 {
		goto lab1  //跳转到lab1
	}

	fmt.Println(1);
	fmt.Println(2);
	
	lab1:
	//执行
	fmt.Println(11);
	fmt.Println(22);

	//output: 11 22
}