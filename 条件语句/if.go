//条件判断语句 if和switch

package main 

import "fmt"

func main() {

	//声明并初始化变量
	i := 1

	// if单支语句
	if i >= 1 {
		fmt.Println("i大于等于1"); //output: i大于等于1
	}

	// if多支语句
	if i < 0 {
		fmt.Println("i小于0");
	} else if i >= 1 && i < 10 {
		fmt.Println("i大于等于1");  //output: i大于等于1
	} else {
		fmt.Println("go.....");
	}
}