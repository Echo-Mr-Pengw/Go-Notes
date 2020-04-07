//算数运算符

package main 

import "fmt"

func main() {

	//声明并初始化变量
	i := 1
	j := 2

	//两个数值的相加
	fmt.Println(i + j)  //output:3
	
	//两个字符串通过"+"相加 表示字符串的连接
	fmt.Println("hello " + "world")  //output: hello world

	//相减
	fmt.Println(i - j, j - i)  //output: -1 1

	//相乘
	fmt.Println(i * j)  //output: 2

	//相除
	fmt.Println(i / j)      //output: 0  两个整数相除 最后为整数，小数部分会被去除
	fmt.Println(1.2 / 0.5)  //output: 2.4

	//自增，只要后++,，没有前++。
	//++i这种是不存在的
	//i++必须单独在一行，否则报错
	i++
	fmt.Println(i)  //output： 2 

	//自减，只要后--,，没有前--。
	//--i这种是不存在的
	//i--必须单独在一行，否则报错
	j--
	fmt.Println(j)  //output： 1 
}