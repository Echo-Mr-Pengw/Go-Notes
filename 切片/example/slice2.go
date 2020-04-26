// 切片的遍历

package main 

import "fmt"

func main() {

	s1 := []int{1, 2, 3}

	//第一种方法
	sLen := len(s1)
	for i := 0; i < sLen; i++ {
		fmt.Println(s1[i]) // output: 1 2 3
	}

	//第二种方法
	for i, v := range s1 {
		fmt.Println(i, v)   
		// output:
		// 0 1
		// 1 2
		// 2 3
	}
}