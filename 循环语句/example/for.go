// for循环语句

package main 

import (
	"fmt"
)

func main() {

	// for循环语句的第一种写法
	for i := 1; i < 10; i++ {
		fmt.Println(i)   //output： 1 2 3 4 5 6 7 8 9
	}

	// for循环语句的第二种写法
	j := 1
	for j < 10 {
		fmt.Println(j)   //output： 1 2 3 4 5 6 7 8 9
		j++
	}

	// for循环语句的第三种写法
	k := 1
	for true {
		if k < 10 {
			fmt.Println(k)  //output： 1 2 3 4 5 6 7 8 9
		} else {
			break
		}

		k++
	}

	// for循环内的局部变量会被重复利用
	a := [2]int{1, 2}
	for i, v := range a {
		fmt.Println(&i, &v)   // &i &v的地址是相同的
	}
}