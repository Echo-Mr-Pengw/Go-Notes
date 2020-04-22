// 变量的重复声明，则报错

package main

import (
	"fmt"
)

func main() {

	var i = 100
	i := 100;       //.\variable5.go:12:4: no new variables on left side of :=
	fmt.Println(i);
}