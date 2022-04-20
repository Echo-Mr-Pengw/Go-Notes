// goto 标签
// 用于跳出多层循环等场景

package main

import "fmt"

func main() {

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				fmt.Printf("i=%v, j=%v \n", i, j) // output: i=0, j=1
				goto exitFor
			}
		}
	}
	exitFor: // 注意: 只能放在下面, 与11行标签配对
	fmt.Println("exit")  // output: exit
}
