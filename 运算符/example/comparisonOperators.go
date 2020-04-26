// 比较运算符，返回的值都是布尔型 true或者false

package main 

import (
	"fmt"
)

func main() {

	// 声明并初始化变量
	i := 1
	j := 2

	//等于==
	fmt.Println(i == j);    //false

	// 不等于 !=
	fmt.Println(i != j);    //true

	// 大于及大于等于 > >=
	fmt.Println(i > j, i >= j);   //false   false
 
	// 小于及小于等于
	fmt.Println(i < j, i <= j);   //true   true

}