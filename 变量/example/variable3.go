// 全局变量声明

package main 

import (
	"fmt"
)

// 声明全局变量方式一
var i = 1
var j = 2

// 声明全局变量方式二
var (
	i1 = 11
	j1 = 22
)

func main() {

	fmt.Println(i, j);
	
	fmt.Println(i1, j1);
}